package reader

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestNewReader(t *testing.T) {
	getReader := func(t *testing.T) io.ReadCloser {
		b, err := os.ReadFile("testdata/label.xml")
		require.NoError(t, err)
		return io.NopCloser(bytes.NewBuffer(b))
	}

	t.Run("will not stream anything when failed to parse", func(t *testing.T) {
		r := getReader(t)
		count := 0
		observable := NewReader[string](context.Background(), r, "test")

		for item := range observable.Observe() {
			require.NoError(t, item.E)
			count++
		}

		require.Zero(t, count)
	})

	type XmlLabel struct {
		ID          int32   `xml:"id"`
		Name        *string `xml:"name"`
		ContactInfo *string `xml:"contactinfo"`
		Profile     *string `xml:"profile"`
		DataQuality *string `xml:"data_quality"`
	}

	t.Run("will parse items when settings are valid", func(t *testing.T) {
		var (
			count = 0
			r     = getReader(t)
			ctx   = context.Background()
		)
		observable := NewReader[XmlLabel](ctx, r, "label")
		for item := range observable.Observe() {
			count++
			require.NoError(t, item.E)
			require.NotNil(t, item.V)
			l := item.V.(*XmlLabel)
			require.Greater(t, l.ID, int32(0))
		}
		require.Equal(t, 5, count)
	})
}
