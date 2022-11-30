package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getTypes(t *testing.T) {
	t.Run("GetPluralTypes must plurals", func(t *testing.T) {
		items := []string{"artist", "release", "label", "master"}
		result := GetPluralTypes(items)
		assert.Contains(t, result, "artists")
		assert.Contains(t, result, "releases")
		assert.Contains(t, result, "masters")
		assert.Contains(t, result, "labels")
		assert.Len(t, result, 4)
	})
	t.Run("GetPluralTypes must filter unrecognized types", func(t *testing.T) {
		items := []string{"artist", "test", "master", "yugiho"}
		result := GetPluralTypes(items)
		assert.Contains(t, result, "artists")
		assert.Contains(t, result, "masters")
		assert.Len(t, result, 2)
	})
	t.Run("GetPluralTypes must return all types when empty or null", func(t *testing.T) {
		items := []string{}
		result := GetPluralTypes(items)
		assert.Contains(t, result, "artists")
		assert.Contains(t, result, "releases")
		assert.Contains(t, result, "masters")
		assert.Contains(t, result, "labels")
		assert.Len(t, result, 4)

		result = GetPluralTypes(nil)
		assert.Contains(t, result, "artists")
		assert.Contains(t, result, "releases")
		assert.Contains(t, result, "masters")
		assert.Contains(t, result, "labels")
		assert.Len(t, result, 4)
	})
}
