// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLabelURL = "label_url"

// LabelURL mapped from table <label_url>
type LabelURL struct {
	LabelID int32  `gorm:"column:label_id;type:integer;primaryKey" json:"label_id"`
	URLHash int64  `gorm:"column:url_hash;type:bigint;primaryKey" json:"url_hash"` // fnv32 encoded hash from url
	URL     string `gorm:"column:url;type:character varying(2048);not null" json:"url"`
}

// TableName LabelURL's table name
func (*LabelURL) TableName() string {
	return TableNameLabelURL
}