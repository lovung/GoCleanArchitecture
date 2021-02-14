package registry

import (
	"github.com/lovung/GoCleanArchitecture/app/internal/domain/repository"
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/persistence/rdbms/gormrepo"

	"github.com/google/wire"
)

// Dependency Injection: All repository set for wire generate
var (
	repositorySet = wire.NewSet(
		gormrepo.NewUserRepository,
		wire.Bind(new(repository.UserRepository), new(*gormrepo.UserRepository)),
	)
)
