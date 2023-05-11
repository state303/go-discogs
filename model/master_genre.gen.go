// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/state303/go-discogs/src/reference"
	"github.com/state303/go-discogs/src/types"
)

const TableNameMasterGenre = "master_genre"

// MasterGenre mapped from table <master_genre>
type MasterGenre struct {
	MasterID int32 `gorm:"column:master_id;type:integer;primaryKey" json:"master_id"`
	GenreID  int32 `gorm:"column:genre_id;type:integer;primaryKey" json:"genre_id"`
}

func (g *MasterGenre) References() []reference.Reference {
	return []reference.Reference{
		&Reference{g.GenreID, types.GENRE},
	}
}

// TableName MasterGenre's table name
func (*MasterGenre) TableName() string {
	return TableNameMasterGenre
}