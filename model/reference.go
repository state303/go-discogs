package model

import (
	"github.com/state303/go-discogs/src/types"
)

type Reference struct {
	ID  int32
	TYP types.Model
}

func (r *Reference) GetID() int32 {
	return r.ID
}

func (r *Reference) GetType() types.Model {
	return r.TYP
}
