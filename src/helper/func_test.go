package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSliceMapper(t *testing.T) {
	t.Run("slice mapper must return items into slice", func(t *testing.T) {
		s := "test string"
		mapper := SliceMapper[string]()
		v, err := mapper(nil, s)
		require.NoError(t, err)
		var forTyp []string
		require.IsType(t, forTyp, v)
		res := v.([]string)
		require.Len(t, res, 1)
		require.Contains(t, res, s)
	})
}

func TestSliceReducer(t *testing.T) {
	t.Run("slice reducer must return first item", func(t *testing.T) {
		reducer := SliceReducer[string]()
		i, err := reducer(nil, nil, []string{"test"})
		require.NoError(t, err)
		require.IsType(t, []string{}, i)
		res := i.([]string)
		require.Contains(t, res, "test")
		require.Len(t, res, 1)
	})

	t.Run("slice reducer must return all items", func(t *testing.T) {
		reducer := SliceReducer[int]()
		i, err := reducer(nil, []int{1, 2, 3}, []int{4})
		require.NoError(t, err)
		require.IsType(t, []int{}, i)
		res := i.([]int)
		require.Len(t, res, 4)
		for _, j := range []int{1, 2, 3, 4} {
			require.Contains(t, res, j)
		}
	})
}
