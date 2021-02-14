package middleware

import (
	"github.com/lovung/GoCleanArchitecture/app/internal/transaction"
	"github.com/lovung/GoCleanArchitecture/pkg/logger"

	"github.com/gin-gonic/gin"
)

// TransactionMiddleware middleware to help manage the transaction
type TransactionMiddleware struct {
	manager transaction.Manager
}

// NewTransactionMiddleware contructor
func NewTransactionMiddleware(manager transaction.Manager) TransactionMiddleware {
	return TransactionMiddleware{
		manager: manager,
	}
}

// StartRequest start the transaction in the beginning of a request
func (mw *TransactionMiddleware) StartRequest(ctx *gin.Context) {
	txn := mw.manager.TxnBegin(ctx)

	logger.Printf("*gorm.DB address: %p", txn)

	ctx.Set(transaction.ContextKey.String(), txn)
	ctx.Next()
}

// EndRequest get error to check if need to commit or rollback
func (mw *TransactionMiddleware) EndRequest(ctx *gin.Context) {
	ctx.Next()

	logger.Debug("run EndRequest")
	err, ok := ctx.Get("error_context_key")
	if !ok {
		mw.manager.TxnCommit(ctx)
	}
	if p := recover(); p != nil {
		logger.Error("found p and rollback ", p)
		mw.manager.TxnRollback(ctx)
	} else if err != nil {
		logger.Debugf("found e and rollback %v", err)
		mw.manager.TxnRollback(ctx)
	} else {
		logger.Debugf("commit transaction %p", mw.manager.GetTxn(ctx))
		mw.manager.TxnCommit(ctx)
	}
}
