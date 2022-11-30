package types

import (
	"regexp"
)

type Model uint8

const (
	ARTIST Model = 1 + iota
	LABEL
	MASTER
	RELEASE
	STYLE
	GENRE
)

const (
	artists  = "artists"
	labels   = "labels"
	masters  = "masters"
	releases = "releases"
)

var (
	pluralTypes = []string{artists, labels, masters, releases}
	typesRegexp = regexp.MustCompile("^artists|labels|masters|releases$")
)

func GetPluralTypes(types []string) []string {
	if len(types) == 0 {
		return pluralTypes
	}

	res := map[string]struct{}{}
	for _, t := range types {
		if t[len(t)-1:] != "s" {
			t += "s"
		}
		if typesRegexp.MatchString(t) {
			res[t] = struct{}{}
		}
	}
	keys := make([]string, 0)
	for s := range res {
		keys = append(keys, s)
	}
	return keys
}
