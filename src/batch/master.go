package batch

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/state303/go-discogs/model"
	"github.com/state303/go-discogs/src/helper"
	"github.com/state303/go-discogs/src/reader"
	"github.com/state303/go-discogs/src/result"
	"gorm.io/gorm"
	"sync"
)

//TODO: add master release step for future use

func GetMasterStep(order Order) Step {
	return func() result.Result {
		updated := 0
		res := UpdateGenreStyle(order, "master")
		updated += res.Count()
		if res.IsErr() {
			return result.NewResult(updated, res.Err())
		}
		res = InsertMaster(order)
		updated += res.Count()
		if res.IsErr() {
			return result.NewResult(updated, res.Err())
		}
		res = InsertMasterRelations(order)
		updated += res.Count()
		if res.IsErr() {
			return result.NewResult(updated, res.Err())
		}
		return result.NewResult(updated, nil)
	}
}

func InsertMaster(order Order) result.Result {
	return InsertSimple[XmlMaster, model.Master](order, "masters", "master")
}

func UpdateGenreStyle(order Order, localName string) result.Result {
	r := newReadCloser(order.getFilePath(), "update genres and styles")

	styles, genres := make(map[string]struct{}), make(map[string]struct{})
	for item := range reader.NewReader[XmlGenreStyle](order.getContext(), r, localName).Observe() {
		gs := item.V.(*XmlGenreStyle)
		for _, s := range gs.Styles {
			styles[s] = struct{}{}
		}
		for _, g := range gs.Genres {
			genres[g] = struct{}{}
		}
	}

	ms, mg := make([]*model.Style, 0), make([]*model.Genre, 0)
	for s := range styles {
		ms = append(ms, &model.Style{Name: s})
	}
	for g := range genres {
		mg = append(mg, &model.Genre{Name: g})
	}

	fmt.Printf("\nScanned %+v styles and %+v genres.\n", len(ms), len(mg))

	tx := order.getDB().
		Session(&gorm.Session{}).
		Clauses(styleConstraint).
		Create(&ms)
	updated := int(tx.RowsAffected)
	fmt.Printf("Updated %+v styles\n", tx.RowsAffected)
	if tx.Error != nil {
		logrus.Errorf("error during styles insertion: %+v\n", tx.Error)
	}
	tx = order.getDB().
		Session(&gorm.Session{}).
		Clauses(genreConstraint).
		Create(&mg)
	updated += int(tx.RowsAffected)
	fmt.Printf("Updated %+v genres\n", tx.RowsAffected)
	if tx.Error != nil {
		logrus.Errorf("error dusing genres insertion: %+v\n", tx.Error)
	}

	fetchedS, fetchedG := make([]*model.Style, 0), make([]*model.Genre, 0)
	order.getDB().Session(&gorm.Session{}).Find(&fetchedS)
	order.getDB().Session(&gorm.Session{}).Find(&fetchedG)

	for _, s := range fetchedS {
		StyleCache.Store(s.Name, s.ID)
	}
	for _, g := range fetchedG {
		GenreCache.Store(g.Name, g.ID)
	}
	fmt.Printf("Cached %+v styles and %+v genres.\n", len(ms), len(mg))
	return result.NewResult(updated, nil)
}

func InsertMasterRelations(order Order) result.Result {
	r := newReadCloser(order.getFilePath(), "updating master relations...")
	var (
		wg   = new(sync.WaitGroup)
		res  = make(chan result.Result)
		done = make(chan struct{}, 1)
	)

	go func() {
		<-reader.NewReader[XmlMasterRelation](order.getContext(), r, "master").
			WindowWithCount(order.getChunkSize()).
			Map(helper.MapWindowedSlice[*XmlMasterRelation]()).
			ForEach(
				WriteMasterRelations(order, res, wg),
				printError(),
				signalDone(done, wg))
	}()

	go func() {
		<-done
		close(res)
	}()

	sum := result.NewResult(0, nil)
	for next := range res {
		sum = sum.Sum(next)
	}

	fmt.Printf("\nUpdated %+v master relations\n", sum.Count())
	return sum
}

func WriteMasterRelations(order Order, res chan result.Result, wg *sync.WaitGroup) func(i interface{}) {
	return func(i interface{}) {
		wg.Add(1) // process takes time, hence add lock scenario
		mrs := i.([]*XmlMasterRelation)
		mv := make([]*model.MasterVideo, 0)
		ms := make([]*model.MasterStyle, 0)
		mg := make([]*model.MasterGenre, 0)
		ma := make([]*model.MasterArtist, 0)
		for _, mr := range mrs {
			if mr == nil {
				continue
			}
			ms = append(ms, mr.GetMasterStyles()...)
			mg = append(mg, mr.GetMasterGenres()...)
			mv = append(mv, mr.GetMasterVideos()...)
			ma = append(ma, mr.GetMasterArtists()...)
		}
		go func(res chan result.Result) {
			defer wg.Done()
			res <- writeThenReport(order, wg, mv, ms, mg, ma)
		}(res)
	}
}
