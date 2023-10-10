package nettesting

import (
	"net/http"
	"net/http/httputil"
	"sync"
)

type TB interface {
	Logf(format string, args ...any)
}

type HTTPProxyHandler struct {
	handlerDone sync.WaitGroup
	hook        func(r *http.Request) bool
	httpProxy   httputil.ReverseProxy
	t           TB
}

func NewHTTPProxyHandler(t TB, hook func(req *http.Request) bool) *HTTPProxyHandler {
	panic("stub")
}

func (h *HTTPProxyHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	panic("stub")
}

func (h *HTTPProxyHandler) Wait() {
	panic("stub")
}

type Embedme interface{}
