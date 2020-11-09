package registry

import (
	"github.com/google/wire"
	"github.com/lovung/GoCleanArchitecture/app/domain/repository"
	"github.com/lovung/GoCleanArchitecture/app/interface/persistence/rdbms/gorm"
)

// Dependency Injection: All repository set for wire generate
var (
	repositorySet = wire.NewSet(
		gorm.NewTxDataSQL,
		gorm.NewSampleRepository,
		wire.Bind(new(repository.SampleRepository), new(*gorm.SampleRepository)),
	)
)
