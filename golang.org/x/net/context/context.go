package context

import (
	"context"
	"time"
)

func Background() Context {
	panic("stub")
}

func TODO() Context {
	panic("stub")
}

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	panic("stub")
}

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	panic("stub")
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	panic("stub")
}

func WithValue(parent Context, key interface{}, val interface{}) Context {
	panic("stub")
}

type Context context.Context

type CancelFunc context.CancelFunc

type Embedme interface{}
