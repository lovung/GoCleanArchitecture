package appctx

import (
	"context"
	"strconv"
)

// Read more: https://blog.golang.org/context#TOC_3.2.
// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

func (k key) String() string {
	return strconv.Itoa(int(k))
}

// Context key constants for responses
// Its value of zero is arbitrary. If this package defined other context keys, they would have
// different integer values.
const (
	MetaContextKey key = iota + 1
	DataContextKey
	ErrorContextKey
	TransactionContextKey
)

// SetValue wrapped the context.WithValue with appctx keys
func SetValue(ctx context.Context, key key, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

// GetValue from app context with key
func GetValue(ctx context.Context, key key) interface{} {
	return ctx.Value(key)
}
