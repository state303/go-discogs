// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDatum = "data"

// Data mapped from table <data>
type Data struct {
	Etag        string    `gorm:"column:etag;type:character varying(200);primaryKey" json:"etag"`                    // ETag representing this data being unique. Used for updating idcache.
	GeneratedAt time.Time `gorm:"column:generated_at;type:timestamp without time zone;not null" json:"generated_at"` // Date this data is generated at.
	Checksum    string    `gorm:"column:checksum;type:character varying(200);not null" json:"checksum"`              // Checksum to validate when fetching dump source.
	/*
		Type of data. Referred as artists, labels, masters, release.
		Always uppercase.
	*/
	TargetType string `gorm:"column:target_type;type:character varying(20);not null" json:"target_type"`
	URI        string `gorm:"column:uri;type:character varying(2048);not null" json:"uri"` // URI to download dump data file.
}

// TableName Data's table name
func (*Data) TableName() string {
	return TableNameDatum
}
