package xsrftoken

import "time"

func Generate(key, userID, actionID string) string {
	panic("stub")
}

func Valid(token, key, userID, actionID string) bool {
	panic("stub")
}

func ValidFor(token, key, userID, actionID string, timeout time.Duration) bool {
	panic("stub")
}

type Embedme interface{}
