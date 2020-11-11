package gorm

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/lovung/GoCleanArchitecture/app/domain/entity"
	"github.com/lovung/GoCleanArchitecture/app/interface/persistence/rdbms/gorm/model"
	"gorm.io/gorm"
)

type SampleRepository struct {
	ctx context.Context
	db  *gorm.DB
}

func NewSampleRepository(
	ctx context.Context,
	db *gorm.DB,
) *SampleRepository {
	return &SampleRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r *SampleRepository) Create(ent entity.Sample) (created entity.Sample, err error) {
	var mdl model.SampleModel

	err = copier.Copy(&mdl, &ent)
	if err != nil {
		return entity.Sample{}, err
	}

	err = r.db.Create(&mdl).Error
	if err != nil {
		return entity.Sample{}, err
	}

	err = copier.Copy(&created, &mdl)
	if err != nil {
		return entity.Sample{}, err
	}
	return created, nil
}

func (r *SampleRepository) GetByID(id interface{}) (ent entity.Sample, err error) {
	var mdl model.SampleModel

	err = r.db.Take(&mdl, id).Error
	if err != nil {
		return entity.Sample{}, err
	}

	err = copier.Copy(&ent, &mdl)
	if err != nil {
		return entity.Sample{}, err
	}
	return ent, nil
}
