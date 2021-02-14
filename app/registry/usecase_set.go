package registry

import (
	"github.com/lovung/GoCleanArchitecture/app/internal/usecase"
	"github.com/lovung/GoCleanArchitecture/app/internal/usecase/interactor"

	"github.com/google/wire"
)

var (
	useCaseSet = wire.NewSet(
		interactor.NewUserUseCase,
		wire.Bind(new(usecase.UserUseCase), new(*interactor.UserUseCase)),
	)
)
