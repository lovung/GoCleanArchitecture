package transaction

import "context"

//go:generate mockgen -destination=./mocktrans/mock_$GOFILE -source=$GOFILE -package=mocktrans

// ContextKey key to save the transaction object
const ContextKey = "transaction_context_key"

// Manager represents operations needed for transaction support.
// It only needs to be implemented once for each database
type Manager interface {
	TxnBegin(ctx context.Context) interface{}
	TxnCommit(ctx context.Context) error
	TxnRollback(ctx context.Context) error
	GetTxn(ctx context.Context) interface{}
}
