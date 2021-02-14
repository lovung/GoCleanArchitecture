package registry

import (
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/handler"

	"github.com/google/wire"
)

//nolint:deadcode,varcheck,unused
var (
	authHanlderSet = wire.NewSet(
		singletonSet,
		repositorySet,
		useCaseSet,
		handler.NewAuthHandler,
	)
)
