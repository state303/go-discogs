package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsKey(t *testing.T) {
	m := make(map[string]string)
	m["hello"] = "world"

	t.Run("must return true", func(t *testing.T) {
		assert.True(t, containsKey(m, "hello"))
	})
	t.Run("must return false", func(t *testing.T) {
		assert.False(t, containsKey(m, "world"))
	})
	t.Run("must return false if nil map", func(t *testing.T) {
		var mm map[string]int
		assert.False(t, containsKey(mm, "hello"))
	})
}

func Test_getKeys(t *testing.T) {
	m := make(map[string]struct{})
	m["test"] = struct{}{}
	m["map"] = struct{}{}
	m["select"] = struct{}{}
	t.Run("get keys must return all", func(t *testing.T) {
		r := getKeys(m)
		assert.Contains(t, r, "test")
		assert.Contains(t, r, "map")
		assert.Contains(t, r, "select")
		assert.Len(t, r, 3)
	})
}

func Test_getTypes(t *testing.T) {
	t.Run("getTypes must plurals", func(t *testing.T) {
		items := []string{"artist", "release", "label", "master"}
		result := getTypes(items)
		assert.Contains(t, result, "artists")
		assert.Contains(t, result, "releases")
		assert.Contains(t, result, "masters")
		assert.Contains(t, result, "labels")
		assert.Len(t, result, 4)
	})
	t.Run("getTypes must filter unrecognized types", func(t *testing.T) {
		items := []string{"artist", "test", "master", "yugiho"}
		result := getTypes(items)
		assert.Contains(t, result, "artists")
		assert.Contains(t, result, "masters")
		assert.Len(t, result, 2)
	})
	t.Run("getTypes must return all types when empty or nil", func(t *testing.T) {
		items := []string{}
		result := getTypes(items)
		assert.Contains(t, result, "artists")
		assert.Contains(t, result, "releases")
		assert.Contains(t, result, "masters")
		assert.Contains(t, result, "labels")
		assert.Len(t, result, 4)

		result = getTypes(nil)
		assert.Contains(t, result, "artists")
		assert.Contains(t, result, "releases")
		assert.Contains(t, result, "masters")
		assert.Contains(t, result, "labels")
		assert.Len(t, result, 4)
	})
}
