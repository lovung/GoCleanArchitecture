package repository

import (
	"context"

	"github.com/lovung/GoCleanArchitecture/app/internal/domain/entity"
)

//go:generate mockgen -destination=./mockrepo/mock_$GOFILE -source=$GOFILE -package=mockrepo

// UserRepository represents all methods which touch the `User` entity
type UserRepository interface {
	Create(ctx context.Context, ent entity.User) (created entity.User, err error)
	GetByID(ctx context.Context, id interface{}) (ent entity.User, err error)
}
