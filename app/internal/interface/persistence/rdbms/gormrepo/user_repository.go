package gormrepo

import (
	"context"

	"github.com/lovung/GoCleanArchitecture/app/internal/domain/entity"
)

// UserRepository repository struct implement the User Repository interface
type UserRepository struct {
	baseRepository
}

// NewUserRepository creates the new instance of User Repository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create method to call DB to create the new user
func (r *UserRepository) Create(ctx context.Context, ent entity.User) (created entity.User, err error) {
	err = ent.GenID()
	if err != nil {
		return
	}
	err = r.DB(ctx).Create(&ent).Error
	if err != nil {
		return
	}
	err = r.DB(ctx).Take(&created, "id = ?", ent.ID).Error
	return created, err
}

// GetByID get the User by ID
func (r *UserRepository) GetByID(ctx context.Context, id interface{}) (ent entity.User, err error) {
	err = r.DB(ctx).Take(&ent, id).Error
	return ent, err
}
