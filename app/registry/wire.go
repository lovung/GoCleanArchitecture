// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package registry

import (
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/handler"
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/middleware"

	"github.com/google/wire"
)

// TransactionMiddleware DI for middleware
func TransactionMiddleware() middleware.TransactionMiddleware {
	wire.Build(txnMwSet)
	return middleware.TransactionMiddleware{}
}

// AuthHandler DI for handler
func AuthHandler() handler.AuthHandler {
	wire.Build(authHanlderSet)
	return handler.AuthHandler{}
}
