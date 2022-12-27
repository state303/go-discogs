package unique

import "github.com/mitchellh/hashstructure/v2"

func Slice[T comparable](items []T) []T {
	m := make(map[uint64]struct{})
	r := make([]T, 0)
	for i := range items {
		hash, err := hashstructure.Hash(items[i], hashstructure.FormatV2, nil)
		if err != nil {
			panic(err)
		}
		if _, ok := m[hash]; !ok {
			m[hash] = struct{}{}
			r = append(r, items[i])
		}
	}
	return r
}
