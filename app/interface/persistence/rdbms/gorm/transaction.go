package gorm

import (
	"context"

	"gorm.io/gorm"
)

// TxDataSQL manages the transaction by implementing the Transaction Manager interface
type TxDataSQL struct {
	db     *gorm.DB
	isDone bool
}

// NewTxDataSQL is the contructor function
func NewTxDataSQL(ctx context.Context, db *gorm.DB) *TxDataSQL {
	return &TxDataSQL{
		db:     db.WithContext(ctx),
		isDone: false,
	}
}

// TxBegin begin a new transaction
func (tds *TxDataSQL) TxBegin() {
	tds.db = tds.db.Begin()
}

// TxRollback rollback a transaction
func (tds *TxDataSQL) TxRollback() {
	if !tds.isDone {
		tds.db.Rollback()
		tds.isDone = true
	}
}

// TxCommit commit a transaction
func (tds *TxDataSQL) TxCommit() (err error) {
	if !tds.isDone {
		err = tds.db.Commit().Error
		tds.isDone = true
	}
	return err
}

// GetTx to get the current transaction of this service
func (tds *TxDataSQL) GetTx() interface{} {
	return tds.db
}
