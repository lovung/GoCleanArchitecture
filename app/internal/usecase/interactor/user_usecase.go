package interactor

import (
	"context"

	"github.com/lovung/GoCleanArchitecture/app/internal/domain/entity"
	"github.com/lovung/GoCleanArchitecture/app/internal/domain/repository"
	"github.com/lovung/GoCleanArchitecture/app/internal/usecase/dto"
	"github.com/lovung/GoCleanArchitecture/pkg/copier"
)

// UserUseCase implements the user use case interface
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase create the new use case
func NewUserUseCase(
	userRepo repository.UserRepository,
) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// Create method to create a new user
func (u *UserUseCase) Create(ctx context.Context, candidate dto.CreateUserRequest) (created dto.OneUserResponse, err error) {
	var ent entity.User
	copier.MustCopy(&ent, &candidate)

	err = ent.HashPassword()
	if err != nil {
		return created, err
	}

	createdEnt, err := u.userRepo.Create(ctx, ent)
	if err != nil {
		return created, err
	}

	copier.MustCopy(&created, &createdEnt)
	return created, nil
}

// GetByID to get the user by ID
func (u *UserUseCase) GetByID(ctx context.Context, id interface{}) (exist dto.OneUserResponse, err error) {
	existEnt, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return dto.OneUserResponse{}, err
	}

	// Do some specific business logic code here
	copier.MustCopy(&exist, &existEnt)
	return exist, nil
}
