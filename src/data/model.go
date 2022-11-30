package data

import "time"

type Data struct {
	ETag        string `xml:"ETag" gorm:"column:etag"`
	GeneratedAt time.Time
	Checksum    string
	TargetType  string `gorm:"column:target_type"`
	Uri         string `xml:"Key"`
}

func (Data) TableName() string {
	return "data"
}
