package transaction

import (
	"context"
	"strconv"
)

//go:generate mockgen -destination=./mocktrans/mock_$GOFILE -source=$GOFILE -package=mocktrans

// Read more: https://blog.golang.org/context#TOC_3.2.
// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

func (k key) String() string {
	return strconv.Itoa(int(k))
}

// ContextKey key to save the transaction object
// Its value of zero is arbitrary. If this package defined other context keys, they would have
// different integer values.
const ContextKey key = 1

// Manager represents operations needed for transaction support.
// It only needs to be implemented once for each database
type Manager interface {
	TxnBegin(ctx context.Context) interface{}
	TxnCommit(ctx context.Context) error
	TxnRollback(ctx context.Context) error
	GetTxn(ctx context.Context) interface{}
}
