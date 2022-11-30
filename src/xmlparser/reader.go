package xmlparser

import (
	"context"
	"github.com/reactivex/rxgo/v2"
)

// ParseItems reads annotated item of types T with context.Context.
// TokenParseOrder will dictate current io.Reader position and io.Closer for resource
func ParseItems[T any](ctx context.Context, order TokenParseOrder) rxgo.Observable {
	c := make(chan rxgo.Item)
	go func(c chan rxgo.Item) {
		parser := NewParser[T]()
		defer func() { close(c); _ = order.Close() }()
		in := parser.Parse(ctx, order).Observe()
		for {
			select {
			case <-ctx.Done():
				return
			case item, open := <-in:
				if !open {
					return
				}
				if item.Error() {
					c <- rxgo.Error(item.E)
					return
				} else {
					c <- rxgo.Of(item.V)
				}
			}
		}
	}(c)
	return rxgo.FromChannel(c)
}
