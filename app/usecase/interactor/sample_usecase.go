package interactor

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/lovung/GoCleanArchitecture/app/domain/entity"
	"github.com/lovung/GoCleanArchitecture/app/domain/repository"
	"github.com/lovung/GoCleanArchitecture/app/usecase/dto"
)

type SampleUseCase struct {
	ctx        context.Context
	sampleRepo repository.SampleRepository
}

func NewSampleUseCase(
	ctx context.Context,
	sampleRepo repository.SampleRepository,
) *SampleUseCase {
	return &SampleUseCase{
		ctx:        ctx,
		sampleRepo: sampleRepo,
	}
}

func (u *SampleUseCase) Create(candidate dto.CreateSampleRequest) (created dto.OneSampleResponse, err error) {
	var ent entity.Sample

	err = copier.Copy(&ent, &candidate)
	if err != nil {
		return dto.OneSampleResponse{}, err
	}

	// Do some specific business logic code here
	createdEnt, err := u.sampleRepo.Create(ent)
	if err != nil {
		return dto.OneSampleResponse{}, err
	}

	err = copier.Copy(&created, &createdEnt)
	if err != nil {
		return dto.OneSampleResponse{}, err
	}
	return created, nil
}

func (u *SampleUseCase) GetByID(id interface{}) (exist dto.OneSampleResponse, err error) {
	existEnt, err := u.sampleRepo.GetByID(id)
	if err != nil {
		return dto.OneSampleResponse{}, err
	}

	// Do some specific business logic code here
	err = copier.Copy(&exist, &existEnt)
	if err != nil {
		return dto.OneSampleResponse{}, err
	}
	return exist, nil
}
