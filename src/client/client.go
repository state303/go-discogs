package client

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"io"
	"net/http"
)

type Client interface {
	Get(ctx context.Context, uri string) rxgo.Observable
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewClient() Client {
	return &clientImpl{http.DefaultClient}
}

type clientImpl struct {
	wc httpClient
}

func (c *clientImpl) Get(_ context.Context, uri string) rxgo.Observable {
	return rxgo.Just(uri)().
		Map(func(ct context.Context, i interface{}) (interface{}, error) {
			withContext, _ := http.NewRequestWithContext(ct, http.MethodGet, uri, nil)
			if resp, err := c.wc.Do(withContext); err != nil {
				return nil, err
			} else {
				b, _ := io.ReadAll(resp.Body)
				return b, nil
			}
		})
}
