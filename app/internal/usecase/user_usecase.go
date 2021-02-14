package usecase

import (
	"context"

	"github.com/lovung/GoCleanArchitecture/app/internal/usecase/dto"
)

//go:generate mockgen -destination=./mockusecase/mock_$GOFILE -source=$GOFILE -package=mockusecase

// UserUseCase interface represents the methods to interact with the User
type UserUseCase interface {
	Create(ctx context.Context, candidate dto.CreateUserRequest) (created dto.OneUserResponse, err error)
	GetByID(ctx context.Context, id interface{}) (exist dto.OneUserResponse, err error)
}
