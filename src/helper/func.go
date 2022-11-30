package helper

import (
	"context"
	"github.com/reactivex/rxgo/v2"
)

func SliceMapper[T any]() func(_ context.Context, i interface{}) (interface{}, error) {
	return func(ctx context.Context, i interface{}) (interface{}, error) {
		return []T{i.(T)}, nil
	}
}

func SliceReducer[T any]() func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
	return func(ctx context.Context, acc interface{}, elem interface{}) (interface{}, error) {
		if acc == nil {
			return elem, nil
		}
		return append(acc.([]T), elem.([]T)...), nil
	}
}

func MergeCount() func(ctx context.Context, i interface{}, i2 interface{}) (interface{}, error) {
	return func(ctx context.Context, acc interface{}, curr interface{}) (interface{}, error) {
		if acc == nil {
			return curr.(int), nil
		}
		return acc.(int) + curr.(int), nil
	}
}

func MapWindowedSlice[T any]() func(ctx context.Context, i interface{}) (interface{}, error) {
	return func(ctx context.Context, i interface{}) (interface{}, error) {
		items := make([]T, 0)
		for item := range i.(rxgo.Observable).Observe() {
			if item.V == nil {
				continue
			}
			items = append(items, item.V.(T))
		}
		return items, nil
	}
}
