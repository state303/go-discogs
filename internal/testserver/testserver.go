package testserver

import (
	"net/http"
	"net/http/httptest"
	"time"
)

type HttpRequest struct {
	*http.Request
	time time.Time
}

func (r *HttpRequest) GetTime() time.Time {
	return r.time
}

type HandleFunc func(requests []*HttpRequest, http http.ResponseWriter, r *http.Request)

func NewServer(h HandleFunc) *TestServer {
	serverRequests := make([]*HttpRequest, 0)
	handler := &HttpHandler{handleFunc: h}
	server := &TestServer{httptest.NewServer(handler), serverRequests}
	handler.server = server
	return server
}

func NewServerWithStaticResponse(response string) *TestServer {
	return NewServer(func(requests []*HttpRequest, writer http.ResponseWriter, r *http.Request) {
		_, _ = writer.Write([]byte(response))
	})
}

type TestServer struct {
	*httptest.Server
	requests []*HttpRequest
}

func (f *TestServer) Requests() []*HttpRequest {
	return f.requests
}

type HttpHandler struct {
	server     *TestServer
	handleFunc HandleFunc
}

func (h *HttpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.server.requests = append(h.server.Requests(), &HttpRequest{request, time.Now()})
	h.handleFunc(h.server.Requests(), writer, request)
}
