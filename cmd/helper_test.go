package cmd

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_fixTypesAsPlural(t *testing.T) {
	samples := []string{"artist", "label", "master", "release"}
	t.Run("must return plurals", func(t *testing.T) {
		k := koanf.New(".")
		err := k.Load(rawbytes.Provider([]byte(`
types:
  - artist
  - label
  - master
  - release
`)), yaml.Parser())

		require.NoError(t, err)
		items := k.Strings("types")
		for _, key := range samples {
			require.Contains(t, items, key)
		}
		fixTypesAsPlural(k)
		items = k.Strings("types")
		for _, key := range samples {
			require.NotContains(t, items, key)
		}
	})

	t.Run("must not complain when no types present", func(t *testing.T) {
		k := koanf.New(".")
		require.NotPanics(t, func() { fixTypesAsPlural(k) })
	})

	t.Run("must not set plural when entry is empty string", func(t *testing.T) {
		k := koanf.New(".")
		err := k.Load(rawbytes.Provider([]byte(`
types:
  - " "
`)), yaml.Parser())
		require.NoError(t, err)
		require.True(t, k.Exists("types"))
		items := k.Strings("types")
		require.Len(t, items, 1)
		fmt.Println(items[0])
		require.Len(t, strings.TrimSpace(items[0]), 0)
		fixTypesAsPlural(k)
		items = k.Strings("types")
		require.Len(t, items, 1)
	})
}
