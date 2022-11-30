package reference

import "github.com/state303/go-discogs/src/types"

type Reference interface {
	GetID() int32
	GetType() types.Model
}

type Referable interface {
	References() []Reference
}
