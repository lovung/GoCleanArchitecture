package registry

import (
	"github.com/google/wire"
	"github.com/lovung/GoCleanArchitecture/pkg/gormutil"
)

var (
	singletonSet = wire.NewSet(
		gormutil.GetDB,
	)
)
