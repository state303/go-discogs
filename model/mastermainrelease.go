package model

import (
	"github.com/state303/go-discogs/src/reference"
	"github.com/state303/go-discogs/src/types"
)

// MasterMainRelease mapped from table <master>
type MasterMainRelease struct {
	ID            int32 `gorm:"column:id;type:integer;primaryKey" json:"id"`
	MainReleaseID int32 `gorm:"column:main_release_id;type:integer" json:"main_release_id"`
}

func (m *MasterMainRelease) References() []reference.Reference {
	return []reference.Reference{
		&Reference{m.ID, types.MASTER},
	}
}

// TableName Master's table name
func (*MasterMainRelease) TableName() string {
	return TableNameMaster
}
