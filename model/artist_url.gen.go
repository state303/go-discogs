// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameArtistURL = "artist_url"

// ArtistURL mapped from table <artist_url>
type ArtistURL struct {
	ArtistID int32  `gorm:"column:artist_id;type:integer;primaryKey" json:"artist_id"`
	URLHash  int64  `gorm:"column:url_hash;type:bigint;primaryKey" json:"url_hash"` // fnv32 encoded hash from url
	URL      string `gorm:"column:url;type:character varying(2048);not null" json:"url"`
}

// TableName ArtistURL's table name
func (*ArtistURL) TableName() string {
	return TableNameArtistURL
}