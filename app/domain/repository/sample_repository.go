package repository

import "github.com/lovung/GoCleanArchitecture/app/domain/entity"

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

// SampleRepository represents all methods which touch the `Sample` entity
type SampleRepository interface {
	Create(ent entity.Sample) (created entity.Sample, err error)
	GetByID(id interface{}) (ent entity.Sample, err error)
}
