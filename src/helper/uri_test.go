package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetLastUriSegment(t *testing.T) {
	t.Run("must get last uri segment", func(t *testing.T) {
		s := "test/my/custom/path"
		x := GetLastUriSegment(s)
		require.Equal(t, "path", x)
	})
	t.Run("must resolve empty path", func(t *testing.T) {
		s := ""
		x := GetLastUriSegment(s)
		require.Equal(t, "", x)
	})
	t.Run("must resolve sole path", func(t *testing.T) {
		s := "here"
		x := GetLastUriSegment(s)
		require.Equal(t, "here", x)
	})
}
