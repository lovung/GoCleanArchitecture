package transaction

import (
	"context"
)

// Manager represents operations needed for transaction support.
// It only needs to be implemented once for each database
type Manager interface {
	TxnBegin(ctx context.Context) context.Context
	TxnCommit(ctx context.Context) error
	TxnRollback(ctx context.Context) error
	GetTxn(ctx context.Context) interface{}
}
