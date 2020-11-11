// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package registry

import (
	"context"

	"github.com/google/wire"
	"github.com/lovung/GoCleanArchitecture/app/usecase"
)

// InitializeSampleUseCase DI for use case
func InitializeSampleUseCase(ctx context.Context) usecase.SampleUseCase {
	wire.Build(sampleUseCaseSet)
	return nil
}
