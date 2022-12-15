package batch

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"github.com/state303/go-discogs/model"
	"github.com/state303/go-discogs/src/cache"
	"github.com/state303/go-discogs/src/helper"
	"github.com/state303/go-discogs/src/reader"
	"github.com/state303/go-discogs/src/result"
)

func InsertSimple[F, T any](order Order, topic string, localName string) result.Result {
	r := newReadCloser(order.getFilePath(), fmt.Sprintf("updating %+v...", topic))
	res := <-reader.NewReader[F](order.getContext(), r, localName).
		FlatMap(Transform).
		Map(registerCache).
		WindowWithCount(order.getChunkSize()).
		Map(helper.MapWindowedSlice[*T]()).
		Map(insertBySlice[*T](order)).
		Reduce(helper.MergeCount()).
		Observe(rxgo.WithCPUPool())
	if res.E != nil {
		return result.NewResult(0, res.E)
	}
	fmt.Printf("\nUpdated %+v %+v\n", res.V.(int), topic)
	return result.NewResult(res.V.(int), res.E)
}

func registerCache(_ context.Context, i interface{}) (interface{}, error) {
	if i == nil {
		return i, nil
	}
	switch o := i.(type) {
	case *model.Artist:
		cache.ArtistIDCache.Store(o.ID, struct{}{})
	case *model.Label:
		cache.LabelIDCache.Store(o.ID, struct{}{})
	case *model.Master:
		cache.MasterIDCache.Store(o.ID, struct{}{})
	}
	return i, nil
}
