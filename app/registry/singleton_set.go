package registry

import (
	"github.com/lovung/GoCleanArchitecture/pkg/gormutil"

	"github.com/google/wire"
)

var (
	singletonSet = wire.NewSet(
		gormutil.GetDB,
	)
)
