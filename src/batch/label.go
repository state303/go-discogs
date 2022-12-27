package batch

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/state303/go-discogs/model"
	"github.com/state303/go-discogs/src/cache"
	"github.com/state303/go-discogs/src/helper"
	"github.com/state303/go-discogs/src/reader"
	"github.com/state303/go-discogs/src/result"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
)

//TODO: add label release step for future use

func GetLabelStep(order Order) Step {
	return func() result.Result {
		updated := 0
		res := insertLabels(order)
		updated += res.Count()
		if res.IsErr() {
			return result.NewResult(updated, res.Err())
		}
		res = insertLabelRelations(order)
		updated += res.Count()
		if res.IsErr() {
			return result.NewResult(updated, res.Err())
		}
		return result.NewResult(updated, nil)
	}
}

func insertLabels(order Order) result.Result {
	return InsertSimple[XmlLabel, model.Label](order, "labels", "label")
}

func insertLabelRelations(order Order) result.Result {
	r := newReadCloser(order.getFilePath(), "updating label relations...")

	var (
		wg   = new(sync.WaitGroup)
		res  = make(chan result.Result)
		done = make(chan struct{}, 1)
	)

	go func() {
		<-reader.NewReader[XmlLabelRelation](order.getContext(), r, "label").
			WindowWithCount(order.getChunkSize()).
			Map(helper.MapWindowedSlice[*XmlLabelRelation]()).
			ForEach(
				writeLabelRelations(order, res, wg), // DoOnNext
				printError(),                        // DoOnError
				signalDone(done, wg))                //DoOnComplete
	}()

	go func() { // wait until done called then close res chan
		<-done
		close(res)
	}()

	sum := result.NewResult(0, nil)
	for next := range res {
		sum = sum.Sum(next)
	}

	fmt.Printf("\nUpdated %+v label relations\n", sum.Count())
	return sum
}

func signalDone(done chan<- struct{}, wg *sync.WaitGroup) func() {
	return func() {
		defer close(done)
		wg.Wait()
		done <- struct{}{}
	}
}

func writeLabelRelations(order Order, res chan result.Result, wg *sync.WaitGroup) func(i interface{}) {
	return func(i interface{}) {
		wg.Add(2)
		u := make([]*model.LabelURL, 0)
		lrs := i.([]*XmlLabelRelation)
		for _, lr := range lrs {
			u = append(u, lr.GetUrls()...)
		}
		go func() { defer wg.Done(); res <- updateLabelsParent(lrs, order.getDB()) }()
		go func() { defer wg.Done(); res <- writeThenReport(order, wg, u) }()
	}
}

func updateLabelsParent(labels []*XmlLabelRelation, db *gorm.DB) result.Result {
	lps := make([]*model.Label, 0)
	for _, v := range labels {
		pid := v.GetParentID()
		if pid == nil {
			continue
		}
		if _, ok := cache.LabelIDCache.Load(*pid); ok {
			logrus.Debugf("\nlookup for lable id %+v failed due to missing cache\n", *pid)
			lps = append(lps, &model.Label{ID: v.ID, ParentID: pid})
		}
	}
	tx := db.Session(&gorm.Session{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"parent_id"}),
	}).CreateInBatches(&lps, len(lps))

	return result.NewResult(int(tx.RowsAffected), tx.Error)
}
