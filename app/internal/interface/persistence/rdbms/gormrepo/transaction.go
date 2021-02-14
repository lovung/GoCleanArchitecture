package gormrepo

import (
	"context"

	"github.com/lovung/GoCleanArchitecture/app/internal/transaction"

	"gorm.io/gorm"
)

// TxnDataSQL manages the transaction by implementing the Transaction Manager interface
type TxnDataSQL struct {
	db *gorm.DB
}

// NewTxnDataSQL is the contructor function
func NewTxnDataSQL(db *gorm.DB) *TxnDataSQL {
	return &TxnDataSQL{
		db: db,
	}
}

// TxnBegin begin a new transaction
func (tds *TxnDataSQL) TxnBegin(ctx context.Context) interface{} {
	return tds.db.WithContext(ctx).Begin()
}

// TxnRollback rollback a transaction
func (tds *TxnDataSQL) TxnRollback(ctx context.Context) (err error) {
	return tds.GetTxn(ctx).(*gorm.DB).Rollback().Error
}

// TxnCommit commit a transaction
func (tds *TxnDataSQL) TxnCommit(ctx context.Context) (err error) {
	return tds.GetTxn(ctx).(*gorm.DB).Commit().Error
}

// GetTxn to get the current transaction of this service
func (tds *TxnDataSQL) GetTxn(ctx context.Context) interface{} {
	db, ok := ctx.Value(transaction.ContextKey).(*gorm.DB)
	if !ok {
		panic("assign to *gorm.DB failed")
	}
	return db
}
