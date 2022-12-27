package unique

import (
	"github.com/state303/go-discogs/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getUniqueSlice(t *testing.T) {
	var (
		artistUrls = []*model.ArtistURL{
			{ArtistID: 33175, URLHash: 1061079901, URL: "https://www.instagram.com/legendarydjamar"},
			{ArtistID: 33175, URLHash: 1061079901, URL: "https://www.instagram.com/legendarydjamar"},
			{ArtistID: 33175, URLHash: 1061079901, URL: "https://www.instagram.com/legendarydjamar"},
			{ArtistID: 33175, URLHash: 1061079901, URL: "https://www.instagram.com/legendarydjamar"},
			{ArtistID: 33175, URLHash: 1061079901, URL: "https://www.instagram.com/legendarydjamar"},
			{ArtistID: 33175, URLHash: 1061079901, URL: "https://www.instagram.com/legendarydjamar"},
			{ArtistID: 33175, URLHash: 1061079901, URL: "https://www.instagram.com/legendarydjamar"},
			{ArtistID: 33175, URLHash: 1061079901, URL: "https://www.instagram.com/legendarydjamar"},
			{ArtistID: 33175, URLHash: 1061079901, URL: "https://www.instagram.com/legendarydjamar"},
			{ArtistID: 33175, URLHash: 1061079901, URL: "https://www.instagram.com/legendarydjamar"},
		}
		result = Slice(artistUrls)
	)
	assert.Len(t, result, 2)
}
