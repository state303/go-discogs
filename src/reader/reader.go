package reader

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"github.com/state303/go-discogs/src/xmlparser"
	"io"
)

func NewReader[T any](ctx context.Context, r io.ReadCloser, localName string) rxgo.Observable {
	return xmlparser.ParseItems[T](ctx, xmlparser.SimpleTokenOrder(r, localName))
}
