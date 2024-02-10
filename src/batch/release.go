package batch

import (
	"fmt"
	"github.com/state303/go-discogs/model"
	"github.com/state303/go-discogs/src/cache"
	"github.com/state303/go-discogs/src/helper"
	"github.com/state303/go-discogs/src/reader"
	"github.com/state303/go-discogs/src/result"
	"github.com/state303/go-discogs/src/unique"
	"gorm.io/gorm/clause"
	"strings"
	"sync"
)

// TODO: double check ptr exceptions on reference

// GetReleaseStep returns a set of steps in a form of composite notary.
// This is a convenient func such that reduces code and adds syntactic sugar, but nothing more.
func GetReleaseStep(order Order) Step {
	return func() result.Result {
		return insertReleases(order)
	}
}

func insertReleases(order Order) result.Result {
	r := newReadCloser(order.getFilePath(), "updating releases...")

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
				doInsertReleases(order, res, wg),
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

func doInsertReleases(order Order, res chan result.Result, wg *sync.WaitGroup) func(i interface{}) {
	return func(i interface{}) {
		wg.Add(1)
		rrs := i.([]*XmlReleaseRelation)

		var (
			g   = make([]*model.Genre, 0)
			s   = make([]*model.Style, 0)
			rel = make([]*model.Release, 0)
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
		)

		for _, rr := range rrs {
			if rr == nil {
				continue
			}
			g = append(g, rr.GetGenres()...)
			s = append(s, rr.GetStyles()...)
		}

		order.getDB().
			Clauses(clause.OnConflict{DoNothing: true}).
			CreateInBatches(filterGenres(g), order.getChunkSize())
		order.getDB().
			Clauses(clause.OnConflict{DoNothing: true}).
			CreateInBatches(filterStyles(s), order.getChunkSize())

		var fg []*model.Genre
		var fs []*model.Style

		order.getDB().Find(&fs)
		order.getDB().Find(&fg)

		for _, v := range fs {
			cache.StyleCache.Store(v.Name, v.ID)
		}
		for _, v := range fg {
			cache.GenreCache.Store(v.Name, v.ID)
		}

		for _, rr := range rrs {
			if rr == nil {
				continue
			}
			rel = append(rel, rr.GetRelease())
			ra = append(ra, rr.GetReleaseArtists()...)
			rg = append(rg, rr.GetReleaseGenres()...)
			rs = append(rs, rr.GetReleaseStyles()...)
			rc = append(rc, rr.GetContracts()...)
			rl = append(rl, rr.GetLabels()...)
			rf = append(rf, rr.GetFormats()...)
			ri = append(ri, rr.GetIdentifiers()...)
			rt = append(rt, rr.GetTracks()...)
			rv = append(rv, rr.GetVideos()...)
			rca = append(rca, rr.GetCreditedArtists()...)
		}

		go func(res chan result.Result) {
			defer wg.Done()
			res <- writeThenReport(order, wg, rel, ra, rc, rs, rg, rl, rf, ri, rt, rv, rca)
		}(res)
	}
}

func filterGenres(genres []*model.Genre) []*model.Genre {
	r := make([]*model.Genre, 0)
	for _, v := range unique.Slice(genres) {
		if name := strings.TrimSpace(v.Name); len(name) == 0 {
			continue
		}
		if _, ok := cache.GenreCache.Load(v); !ok {
			r = append(r, v)
		}
	}
	return r
}

func filterStyles(styles []*model.Style) []*model.Style {
	r := make([]*model.Style, 0)
	for _, v := range unique.Slice(styles) {
		if name := strings.TrimSpace(v.Name); len(name) == 0 {
			continue
		}
		if _, ok := cache.StyleCache.Load(v); !ok {
			r = append(r, v)
		}
	}
	return r
}
