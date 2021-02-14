package registry

import (
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/persistence/rdbms/gormrepo"
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/middleware"
	"github.com/lovung/GoCleanArchitecture/app/internal/transaction"

	"github.com/google/wire"
)

//nolint:deadcode,varcheck,unused
var (
	txnMwSet = wire.NewSet(
		singletonSet,
		gormrepo.NewTxnDataSQL,
		wire.Bind(new(transaction.Manager), new(*gormrepo.TxnDataSQL)),
		middleware.NewTransactionMiddleware,
	)
)
