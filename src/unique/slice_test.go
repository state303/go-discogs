package unique

import (
	"github.com/state303/go-discogs/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getUniqueSlice(t *testing.T) {
	var (
		artistUrls = []*model.ArtistURL{
			{33175, 1061079901, "https://www.instagram.com/legendarydjamar"},
			{33175, 1061079901, "https://www.instagram.com/legendarydjamar"},
			{33175, 1061079901, "https://www.instagram.com/legendarydjamar"},
			{33175, 1061079901, "https://www.instagram.com/legendarydjamar"},
			{33175, 1061079901, "https://www.instagram.com/legendarydjamar"},
			{33175, 1061079901, "https://www.instagram.com/legendarydjamar"},
			{33177, 1061079901, "https://www.instagram.com/legendarydjamar"},
		}
		result = Slice(artistUrls)
	)
	assert.Len(t, result, 2)
}
