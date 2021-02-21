package registry

import (
	"github.com/google/wire"
	"github.com/lovung/GoCleanArchitecture/pkg/gormer"
)

var (
	singletonSet = wire.NewSet(
		gormer.GetDB,
	)
)
