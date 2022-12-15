package batch

import (
	"context"
	"github.com/state303/go-discogs/src/reader"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReleaseRead(t *testing.T) {
	var (
		c = context.Background()
		r = newReadCloser("testdata/release.xml.gz", "test-read-release")
		n = "release"
	)
	obs := reader.NewReader[XmlRelease](c, r, n)

	s := make([]*XmlRelease, 0)
	for r := range obs.Observe() {
		if r.V == nil {
			continue
		}
		x := r.V.(*XmlRelease)
		s = append(s, x)
		require.NotNil(t, x.Status)
		require.NotNil(t, x.ListedReleaseDate)
	}

	require.Len(t, s, 3)
	require.True(t, s[0].MasterInfo.IsMaster)
	require.True(t, s[1].MasterInfo.IsMaster)
	require.False(t, s[2].MasterInfo.IsMaster)
}

func TestReleaseRelationRead(t *testing.T) {
	var (
		c = context.Background()
		r = newReadCloser("testdata/release.xml.gz", "test-read-release")
		n = "release"
	)
	obs := reader.NewReader[XmlReleaseRelation](c, r, n)
	s := make([]*XmlReleaseRelation, 0)
	for r := range obs.Observe() {
		require.NoError(t, r.E)
		require.NotNil(t, r.V)
		s = append(s, r.V.(*XmlReleaseRelation))
	}
	require.Len(t, s, 3)
}
