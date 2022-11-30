package batch

import (
	"sync"
)

var (
	// StyleCache stores name and id in form of string and int32
	StyleCache = &sync.Map{}
	// GenreCache stores name and id in form of string and int32
	GenreCache = &sync.Map{}
	// ArtistIDCache stores id in form of int32, struct{}
	ArtistIDCache = &sync.Map{}
	// LabelIDCache stores id in form of int32, struct{}
	LabelIDCache = &sync.Map{}
	// MasterIDCache stores id in form of int32, struct{}
	MasterIDCache = &sync.Map{}
)
