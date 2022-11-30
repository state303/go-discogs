package xmlparser

import "encoding/xml"

type TokenMatcher func(token xml.Token) bool

func NewStartTokenNameMatcher(name string) TokenMatcher {
	return func(token xml.Token) bool {
		if token == nil {
			return false
		}
		if se, ok := token.(xml.StartElement); !ok {
			return false
		} else {
			return se.Name.Local == name
		}
	}
}
