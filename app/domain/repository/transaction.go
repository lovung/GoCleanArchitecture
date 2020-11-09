package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

// TransactionManager represents operations needed for transaction support.
// It only needs to be implemented once for each database
type TransactionManager interface {
	TxBegin()
	TxCommit() error
	TxRollback()
	GetTx() interface{}
}
