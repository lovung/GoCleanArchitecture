package gormrepo

import (
	"context"

	"github.com/lovung/GoCleanArchitecture/app/internal/transaction"

	"gorm.io/gorm"
)

type baseRepository struct{}

// DB to get the transaction to Database from context
func (r *baseRepository) DB(ctx context.Context) *gorm.DB {
	gormDB, ok := ctx.Value(transaction.ContextKey).(*gorm.DB)
	if !ok {
		panic("can not get the gorm.DB in context")
	}
	return gormDB
}
