package registry

import (
	"github.com/google/wire"
	"github.com/lovung/GoCleanArchitecture/app/usecase"
	"github.com/lovung/GoCleanArchitecture/app/usecase/interactor"
)

var (
	sampleUseCaseSet = wire.NewSet(
		singletonSet,
		repositorySet,
		interactor.NewSampleUseCase,
		wire.Bind(new(usecase.SampleUseCase), new(*interactor.SampleUseCase)),
	)
)
