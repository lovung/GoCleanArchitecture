package usecase

import "github.com/lovung/GoCleanArchitecture/app/usecase/dto"

type SampleUseCase interface {
	Create(candidate dto.CreateSampleRequest) (created dto.OneSampleResponse, err error)
	GetByID(id interface{}) (exist dto.OneSampleResponse, err error)
}
