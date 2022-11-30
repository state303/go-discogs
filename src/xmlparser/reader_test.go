package xmlparser

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
	"time"
)

func TestParseItems(t *testing.T) {
	t.Run("must read empty source", func(t *testing.T) {
		b := io.NopCloser(bytes.NewBuffer([]byte("test my string")))

		obs := ParseItems[string](context.Background(), SimpleTokenOrder(b, "hello"))

		count := 0
		for range obs.Observe() {
			count++
		}

		require.Zero(t, count)
	})

	type TestData struct {
		ETag        string `xml:"ETag" gorm:"column:etag"`
		GeneratedAt time.Time
		Checksum    string
		TargetType  string `gorm:"column:target_type"`
		Uri         string `xml:"Key"`
	}

	t.Run("must read target source", func(t *testing.T) {
		f, err := os.Open("testdata/data.xml")
		require.NoError(t, err)
		order := SimpleTokenOrder(f, "Contents")
		observable := ParseItems[TestData](context.Background(), order)
		reads := make([]*TestData, 0)
		for item := range observable.Observe() {
			assert.NoError(t, item.E)
			reads = append(reads, item.V.(*TestData))
		}
		require.NotEmpty(t, reads)
		for _, v := range reads {
			require.NotNil(t, v)
		}
	})

	t.Run("must check capitalized str", func(t *testing.T) {
		f, err := os.Open("testdata/data.xml")
		require.NoError(t, err)
		order := SimpleTokenOrder(f, "contents") // requires Contents, not contents
		observable := ParseItems[TestData](context.Background(), order)
		reads := make([]*TestData, 0)
		for item := range observable.Observe() {
			require.NoError(t, item.E)
			require.NotNil(t, item.V) // last item not being nil
			reads = append(reads, item.V.(*TestData))
		}
		require.Empty(t, reads)
	})

	t.Run("must terminate on context cancellation", func(t *testing.T) {
		f, err := os.Open("testdata/data.xml")
		require.NoError(t, err)
		order := SimpleTokenOrder(f, "Contents")
		observable := ParseItems[TestData](context.Background(), order)
		reads := make([]*TestData, 0)
		for item := range observable.Observe() {
			require.NoError(t, item.E)
			require.NotNil(t, item.V)
			reads = append(reads, item.V.(*TestData))
		}

		f, err = os.Open("testdata/data.xml")
		require.NoError(t, err)
		order = SimpleTokenOrder(f, "Contents")

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			time.AfterFunc(time.Millisecond*2, cancel)
		}()

		observable = ParseItems[TestData](ctx, order)
		nextReads := make([]*TestData, 0)
		for item := range observable.Observe() {
			require.NoError(t, item.E)
			require.NotNil(t, item.V)
			nextReads = append(nextReads, item.V.(*TestData))
		}

		require.Less(t, len(nextReads), len(reads))

		for idx := range nextReads {
			p, c := reads[idx], nextReads[idx]
			// compare, such that parse has been executed in exactly same order
			require.Equal(t, p.ETag, c.ETag)
		}
	})
}
