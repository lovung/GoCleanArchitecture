package middleware

import (
	"github.com/lovung/GoCleanArchitecture/app/internal/appctx"
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
	newCtx := mw.manager.TxnBegin(ctx.Request.Context())
	ctx.Request = ctx.Request.WithContext(newCtx)
	ctx.Next()
}

// EndRequest get error to check if need to commit or rollback
func (mw *TransactionMiddleware) EndRequest(ctx *gin.Context) {
	ctx.Next()
	err := appctx.GetValue(ctx.Request.Context(), appctx.ErrorContextKey)
	if p := recover(); p != nil {
		logger.Error("found p and rollback ", p)
		mw.manager.TxnRollback(ctx.Request.Context())
	} else if err != nil {
		logger.Debugf("found e and rollback %v", err)
		mw.manager.TxnRollback(ctx.Request.Context())
	} else {
		logger.Debugf("commit transaction %p", mw.manager.GetTxn(ctx.Request.Context()))
		mw.manager.TxnCommit(ctx.Request.Context())
	}
}
