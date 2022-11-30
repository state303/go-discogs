package batch

import (
	"github.com/state303/go-discogs/model"
	"github.com/state303/go-discogs/src/result"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
)

func writeThenReport(order Order, wg *sync.WaitGroup, slices ...interface{}) result.Result {
	wg.Add(1)
	defer wg.Done()
	return NewWriter(order.getDB()).Write(order.getChunkSize(), slices...)
}

type Writer interface {
	Write(chunkSize int, items ...interface{}) result.Result
}

type gormWriter struct {
	db *gorm.DB
}

var NewWriter = newWriter

func newWriter(db *gorm.DB) Writer {
	return &gormWriter{db: db}
}

func (g gormWriter) Write(chunkSize int, slices ...interface{}) result.Result {
	var (
		updated = 0
		err     error
	)
	err = nil

	for _, slice := range slices {
		if err != nil {
			break
		}
		var r result.Result
		switch o := slice.(type) {
		case []*model.Artist:
			r = doWrite[*model.Artist](o, chunkSize, g.db)
		case []*model.ArtistURL:
			r = doWrite[*model.ArtistURL](o, chunkSize, g.db)
		case []*model.ArtistAlias:
			r = doWrite[*model.ArtistAlias](o, chunkSize, g.db)
		case []*model.ArtistGroup:
			r = doWrite[*model.ArtistGroup](o, chunkSize, g.db)
		case []*model.ArtistNameVariation:
			r = doWrite[*model.ArtistNameVariation](o, chunkSize, g.db)
		case []*model.Label:
			r = doWrite[*model.Label](o, chunkSize, g.db)
		case []*model.LabelURL:
			r = doWrite[*model.LabelURL](o, chunkSize, g.db)
		case []*model.LabelRelease:
			r = doWrite[*model.LabelRelease](o, chunkSize, g.db)
		case []*model.Master:
			r = doWrite[*model.Master](o, chunkSize, g.db)
		case []*model.MasterMainRelease:
			r = doWrite[*model.MasterMainRelease](o, chunkSize, g.db)
		case []*model.MasterArtist:
			r = doWrite[*model.MasterArtist](o, chunkSize, g.db)
		case []*model.MasterGenre:
			r = doWrite[*model.MasterGenre](o, chunkSize, g.db)
		case []*model.MasterStyle:
			r = doWrite[*model.MasterStyle](o, chunkSize, g.db)
		case []*model.MasterVideo:
			r = doWrite[*model.MasterVideo](o, chunkSize, g.db)
		case []*model.Release:
			r = doWrite[*model.Release](o, chunkSize, g.db)
		case []*model.ReleaseArtist:
			r = doWrite[*model.ReleaseArtist](o, chunkSize, g.db)
		case []*model.ReleaseContract:
			r = doWrite[*model.ReleaseContract](o, chunkSize, g.db)
		case []*model.ReleaseFormat:
			r = doWrite[*model.ReleaseFormat](o, chunkSize, g.db)
		case []*model.ReleaseCreditedArtist:
			r = doWrite[*model.ReleaseCreditedArtist](o, chunkSize, g.db)
		case []*model.ReleaseGenre:
			r = doWrite[*model.ReleaseGenre](o, chunkSize, g.db)
		case []*model.ReleaseStyle:
			r = doWrite[*model.ReleaseStyle](o, chunkSize, g.db)
		case []*model.ReleaseIdentifier:
			r = doWrite[*model.ReleaseIdentifier](o, chunkSize, g.db)
		case []*model.ReleaseImage:
			r = doWrite[*model.ReleaseImage](o, chunkSize, g.db)
		case []*model.ReleaseTrack:
			r = doWrite[*model.ReleaseTrack](o, chunkSize, g.db)
		case []*model.ReleaseVideo:
			r = doWrite[*model.ReleaseVideo](o, chunkSize, g.db)
		}
		if r != nil {
			updated += r.Count()
			err = r.Err()
		}
	}

	return result.NewResult(updated, err)
}

func doWrite[T any](items []T, chunkSize int, db *gorm.DB) result.Result {
	var (
		start     = 0
		end       = chunkSize
		resultSum = result.NewResult(0, nil)
		size      = len(items)
		cl        clause.OnConflict
	)
	if len(items) > 0 {
		cl = ExtractClause(items[0])
	}
	for {
		if start >= size || resultSum.Err() != nil {
			return resultSum
		}
		if end > size {
			end = size
		}
		part := items[start:end]
		tx := db.Clauses(cl).CreateInBatches(&part, len(part))
		resultSum = resultSum.Sum(result.NewResult(int(tx.RowsAffected), tx.Error))
		start += chunkSize
		end += chunkSize
	}
}
