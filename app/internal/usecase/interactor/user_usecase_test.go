package interactor

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/lovung/GoCleanArchitecture/app/internal/domain/entity"
	"github.com/lovung/GoCleanArchitecture/app/internal/domain/repository"
	"github.com/lovung/GoCleanArchitecture/app/internal/domain/repository/mockrepo"
	"github.com/lovung/GoCleanArchitecture/app/internal/usecase/dto"

	"github.com/golang/mock/gomock"
)

func TestUserUseCase_Create(t *testing.T) {
	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		ctx       context.Context
		candidate dto.CreateUserRequest
	}
	type wants struct {
		wantCreated dto.OneUserResponse
		wantErr     bool
	}

	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ctx := context.Background()
	runFunc := func(t *testing.T, fields fields, args args, w wants) {
		u := &UserUseCase{
			userRepo: fields.userRepo,
		}
		gotCreated, err := u.Create(args.ctx, args.candidate)
		if (err != nil) != w.wantErr {
			t.Errorf("UserUseCase.Create() error = %v, wantErr %v", err, w.wantErr)
			return
		}
		if !reflect.DeepEqual(gotCreated, w.wantCreated) {
			t.Errorf("UserUseCase.Create() = %v, want %v", gotCreated, w.wantCreated)
		}
	}

	mUserRepo := mockrepo.NewMockUserRepository(mockCtrl)
	_fields := fields{
		userRepo: mUserRepo,
	}

	t.Run("#1: Create new user failed", func(t *testing.T) {
		mUserRepo.EXPECT().Create(ctx, gomock.Any()).
			Return(entity.User{}, errors.New("error when creating"))
		runFunc(t, _fields, args{
			ctx,
			dto.CreateUserRequest{
				Password: "password",
				Username: "username",
			},
		}, wants{
			dto.OneUserResponse{},
			true,
		})
	})
	t.Run("#2: Success", func(t *testing.T) {
		mUserRepo.EXPECT().Create(ctx, gomock.Any()).
			Return(entity.User{
				ID:       "01ENSP2J4H9QRMWYJ3ZQTYEGYJ",
				Username: "username",
			}, nil)
		runFunc(t, _fields, args{
			ctx,
			dto.CreateUserRequest{
				Password: "password",
				Username: "username",
			},
		}, wants{
			dto.OneUserResponse{
				ID:       "01ENSP2J4H9QRMWYJ3ZQTYEGYJ",
				Username: "username",
			},
			false,
		})
	})
}

func TestNewUserUseCase(t *testing.T) {
	type args struct {
		userRepo repository.UserRepository
	}
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mUserRepo := mockrepo.NewMockUserRepository(mockCtrl)
	tests := []struct {
		name string
		args args
		want *UserUseCase
	}{
		{
			args: args{
				userRepo: mUserRepo,
			},
			want: &UserUseCase{
				userRepo: mUserRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUseCase(tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUseCase_GetByID(t *testing.T) {
	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		ctx context.Context
		id  interface{}
	}
	type wants struct {
		wantExist dto.OneUserResponse
		wantErr   bool
	}
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ctx := context.Background()
	runFunc := func(t *testing.T, fields fields, args args, w wants) {
		u := &UserUseCase{
			userRepo: fields.userRepo,
		}
		gotExist, err := u.GetByID(args.ctx, args.id)
		if (err != nil) != w.wantErr {
			t.Errorf("UserUseCase.GetByID() error = %v, wantErr %v", err, w.wantErr)
			return
		}
		if !reflect.DeepEqual(gotExist, w.wantExist) {
			t.Errorf("UserUseCase.GetByID() = %v, want %v", gotExist, w.wantExist)
		}
	}
	mUserRepo := mockrepo.NewMockUserRepository(mockCtrl)
	_fields := fields{
		userRepo: mUserRepo,
	}

	t.Run("#1: Get user by ID failed", func(t *testing.T) {
		mUserRepo.EXPECT().GetByID(ctx, "01ENSP2J4H9QRMWYJ3ZQTYEGYJ").
			Return(entity.User{}, errors.New("error or not found"))
		runFunc(t, _fields, args{
			ctx,
			"01ENSP2J4H9QRMWYJ3ZQTYEGYJ",
		}, wants{
			dto.OneUserResponse{},
			true,
		})
	})
	t.Run("#2: Found", func(t *testing.T) {
		mUserRepo.EXPECT().GetByID(ctx, "01ENSP2J4H9QRMWYJ3ZQTYEGYJ").
			Return(entity.User{
				ID:       "01ENSP2J4H9QRMWYJ3ZQTYEGYJ",
				Username: "username",
			}, nil)
		runFunc(t, _fields, args{
			ctx,
			"01ENSP2J4H9QRMWYJ3ZQTYEGYJ",
		}, wants{
			dto.OneUserResponse{
				ID:       "01ENSP2J4H9QRMWYJ3ZQTYEGYJ",
				Username: "username",
			},
			false,
		})
	})
}
