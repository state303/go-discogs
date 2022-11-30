package client

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_clientImpl_Get(t *testing.T) {
	t.Run("must handle GET", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			_, _ = writer.Write([]byte("TEST"))
		}))
		defer server.Close()
		c := NewClient()
		res := <-c.Get(context.Background(), server.URL).Observe()
		assert.NoError(t, res.E)
		assert.Equal(t, []byte("TEST"), res.V)
	})
	t.Run("must handle error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Content-Length", "-1")
		}))
		defer server.Close()
		c := &clientImpl{wc: &wc{}}
		res := <-c.Get(context.Background(), server.URL).Observe()
		assert.Error(t, res.E)
	})
}

func TestNewClient(t *testing.T) {
	t.Run("NewClient call", func(t *testing.T) {
		assert.NotNil(t, NewClient())
	})
}

type wc struct{}

func (*wc) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(strings.NewReader(`Internal server error`)),
	}, fmt.Errorf("test")
}
