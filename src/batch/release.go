package batch

import (
	"fmt"
	"github.com/state303/go-discogs/model"
	"github.com/state303/go-discogs/src/helper"
	"github.com/state303/go-discogs/src/reader"
	"github.com/state303/go-discogs/src/result"
	"sync"
)

// TODO: double check ptr exceptions on reference

// GetReleaseStep returns a set of steps in a form of composite notary.
// This is a convenient func such that reduces code and adds syntactic sugar, but nothing more.
func GetReleaseStep(order Order) Step {
	return func() result.Result {
		updated := 0
		res := UpdateGenreStyle(order, "release")
		updated += res.Count()
		if res.IsErr() {
			return result.NewResult(updated, res.Err())
		}
		res = insertReleases(order)
		updated += res.Count()
		if res.IsErr() {
			return result.NewResult(updated, res.Err())
		}
		res = insertReleaseRelations(order)
		updated += res.Count()
		if res.IsErr() {
			return result.NewResult(updated, res.Err())
		}
		return result.NewResult(updated, nil)
	}
}

func insertReleases(order Order) result.Result {
	return InsertSimple[XmlRelease, model.Release](order, "releases", "release")
}

func insertReleaseRelations(order Order) result.Result {
	r := newReadCloser(order.getFilePath(), "updating release relations...")

	var (
		wg   = new(sync.WaitGroup)
		res  = make(chan result.Result)
		done = make(chan struct{}, 1)
	)

	go func() {
		<-reader.NewReader[XmlReleaseRelation](order.getContext(), r, "release").
			WindowWithCount(order.getChunkSize()).
			Map(helper.MapWindowedSlice[*XmlReleaseRelation]()).
			ForEach(
				writeReleaseRelations(order, res, wg),
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

	fmt.Printf("\nUpdated %+v release relations\n", sum.Count())
	return sum
}

func writeReleaseRelations(order Order, res chan result.Result, wg *sync.WaitGroup) func(i interface{}) {
	return func(i interface{}) {
		wg.Add(1)
		rrs := i.([]*XmlReleaseRelation)

		var (
			ra  = make([]*model.ReleaseArtist, 0)
			rca = make([]*model.ReleaseCreditedArtist, 0)
			rc  = make([]*model.ReleaseContract, 0)
			rf  = make([]*model.ReleaseFormat, 0)
			rs  = make([]*model.ReleaseStyle, 0)
			rg  = make([]*model.ReleaseGenre, 0)
			ri  = make([]*model.ReleaseIdentifier, 0)
			rt  = make([]*model.ReleaseTrack, 0)
			rv  = make([]*model.ReleaseVideo, 0)
			rl  = make([]*model.LabelRelease, 0)
			rm  = make([]*model.MasterMainRelease, 0)
		)

		for _, rr := range rrs {
			if rr == nil {
				continue
			}
			ra = append(ra, rr.GetReleaseArtists()...)
			rg = append(rg, rr.GetReleaseGenres()...)
			rs = append(rs, rr.GetReleaseStyles()...)
			rc = append(rc, rr.GetContracts()...)
			rl = append(rl, rr.GetLabels()...)
			rf = append(rf, rr.GetFormats()...)
			ri = append(ri, rr.GetIdentifiers()...)
			rt = append(rt, rr.GetTracks()...)
			rv = append(rv, rr.GetVideos()...)
			rm = append(rm, rr.GetMasterReleases()...)
			rca = append(rca, rr.GetCreditedArtists()...)
		}

		go func(res chan result.Result) {
			defer wg.Done()
			res <- writeThenReport(order, wg, ra, rc, rs, rg, rl, rf, ri, rt, rv, rca, rm)
		}(res)
	}
}
