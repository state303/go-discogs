package helper

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"strings"
)

func SliceMapper[T any]() func(_ context.Context, i interface{}) (interface{}, error) {
	return func(_ context.Context, i interface{}) (interface{}, error) {
		return []T{i.(T)}, nil
	}
}

func SliceReducer[T any]() func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
	return func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
		if acc == nil {
			return elem, nil
		}
		return append(acc.([]T), elem.([]T)...), nil
	}
}

func MergeCount() func(_ context.Context, i interface{}, i2 interface{}) (interface{}, error) {
	return func(_ context.Context, acc interface{}, curr interface{}) (interface{}, error) {
		if acc == nil {
			return curr.(int), nil
		}
		return acc.(int) + curr.(int), nil
	}
}

func MapWindowedSlice[T any]() func(_ context.Context, i interface{}) (interface{}, error) {
	return func(_ context.Context, i interface{}) (interface{}, error) {
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

func FilterStr(s *string) *string {
	if s == nil {
		return nil
	}
	tmp := strings.TrimSpace(*s)
	if len(tmp) == 0 {
		return nil
	} else {
		return &tmp
	}
}
