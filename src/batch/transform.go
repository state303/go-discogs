package batch

import (
	"github.com/reactivex/rxgo/v2"
)

type Transformable interface {
	Transform() rxgo.Observable
}

func Transform(item rxgo.Item) rxgo.Observable {
	if v, ok := item.V.(Transformable); ok {
		return v.Transform()
	}
	return rxgo.Just(item.V)()
}
