package ctxhttp

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

func Do(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	panic("stub")
}

func Get(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	panic("stub")
}

func Head(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	panic("stub")
}

func Post(ctx context.Context, client *http.Client, url string, bodyType string, body io.Reader) (*http.Response, error) {
	panic("stub")
}

func PostForm(ctx context.Context, client *http.Client, url string, data url.Values) (*http.Response, error) {
	panic("stub")
}

type Embedme interface{}
