package gormrepo

import (
	"context"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/lovung/GoCleanArchitecture/app/internal/appctx"
	"github.com/lovung/GoCleanArchitecture/app/internal/domain/entity"
	"github.com/lovung/GoCleanArchitecture/pkg/testhelper"
	"gorm.io/gorm"
)

func TestNewUserRepository(t *testing.T) {
	testCases := []struct {
		name string
		want *UserRepository
	}{
		{
			want: &UserRepository{
				baseRepository: baseRepository{},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Create(t *testing.T) {
	gDB, mock, err := testhelper.OpenDBConnection()
	if err != nil {
		panic(err)
	}
	type fields struct {
		baseRepository baseRepository
	}
	type args struct {
		ctx context.Context
		ent entity.User
	}
	type wants struct {
		wantCreated entity.User
		wantErr     bool
	}
	runFunc := func(t *testing.T, fields fields, args args, w wants) {
		r := &UserRepository{
			baseRepository: fields.baseRepository,
		}
		gotCreated, err := r.Create(args.ctx, args.ent)
		if (err != nil) != w.wantErr {
			t.Errorf("UserRepository.Create() error = %v, wantErr %v", err, w.wantErr)
			return
		}
		if diff := cmp.Diff(gotCreated, w.wantCreated); diff != "" {
			t.Errorf("UserRepository.Create() = %v", diff)
		}
	}
	t.Run("#1: Success", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), appctx.TransactionContextKey, gDB)
		insertQuery := "INSERT INTO `users`(.*) VALUES (.*)"
		mock.ExpectExec(insertQuery).WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

		query := regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? LIMIT 1")
		mock.ExpectQuery(query).WillReturnRows(
			sqlmock.NewRows([]string{"ID", "Username", "Password"}).
				AddRow(
					"01ENSP2J4H9QRMWYJ3ZQTYEGYJ",
					"username",
					"$argon2id$v=19$m=65536,t=8,p=8$rWzlADlkY3v8OsLwVKYFGA$owGB0DwHIFywicD4Ydqz2FY5LT406eNUD7aCgMytqZlmuWnHWkNK3Nl5R+aAhV5tVWOqUObqZaO5+zsGwMNBAA",
				),
		)
		runFunc(t,
			fields{baseRepository{}},
			args{
				ctx: ctx,
				ent: entity.User{
					Username: "username",
					Password: "$argon2id$v=19$m=65536,t=8,p=8$rWzlADlkY3v8OsLwVKYFGA$owGB0DwHIFywicD4Ydqz2FY5LT406eNUD7aCgMytqZlmuWnHWkNK3Nl5R+aAhV5tVWOqUObqZaO5+zsGwMNBAA",
				},
			},
			wants{
				wantCreated: entity.User{
					ID:       "01ENSP2J4H9QRMWYJ3ZQTYEGYJ",
					Username: "username",
					Password: "$argon2id$v=19$m=65536,t=8,p=8$rWzlADlkY3v8OsLwVKYFGA$owGB0DwHIFywicD4Ydqz2FY5LT406eNUD7aCgMytqZlmuWnHWkNK3Nl5R+aAhV5tVWOqUObqZaO5+zsGwMNBAA",
				},
				wantErr: false,
			},
		)
	})
	t.Run("#2: Failed when creating", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), appctx.TransactionContextKey, gDB)
		insertQuery := "INSERT INTO `users`(.*) VALUES (.*)"
		mock.ExpectExec(insertQuery).WillReturnError(gorm.ErrInvalidData)

		runFunc(t,
			fields{baseRepository{}},
			args{
				ctx: ctx,
				ent: entity.User{
					Username: "username",
					Password: "$argon2id$v=19$m=65536,t=8,p=8$rWzlADlkY3v8OsLwVKYFGA$owGB0DwHIFywicD4Ydqz2FY5LT406eNUD7aCgMytqZlmuWnHWkNK3Nl5R+aAhV5tVWOqUObqZaO5+zsGwMNBAA",
				},
			},
			wants{
				wantCreated: entity.User{},
				wantErr:     true,
			},
		)
	})
	t.Run("#3: Failed when select again", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), appctx.TransactionContextKey, gDB)
		insertQuery := "INSERT INTO `users`(.*) VALUES (.*)"
		mock.ExpectExec(insertQuery).WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

		query := regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? LIMIT 1")
		mock.ExpectQuery(query).WillReturnError(gorm.ErrRecordNotFound)
		runFunc(t,
			fields{baseRepository{}},
			args{
				ctx: ctx,
				ent: entity.User{
					Username: "username",
					Password: "$argon2id$v=19$m=65536,t=8,p=8$rWzlADlkY3v8OsLwVKYFGA$owGB0DwHIFywicD4Ydqz2FY5LT406eNUD7aCgMytqZlmuWnHWkNK3Nl5R+aAhV5tVWOqUObqZaO5+zsGwMNBAA",
				},
			},
			wants{
				wantCreated: entity.User{},
				wantErr:     true,
			},
		)
	})
	t.Run("#4: [ShouldPanic] Can not found DB in context", func(t *testing.T) {
		testhelper.ShouldPanic(t, func() {
			ctx := context.Background()
			insertQuery := "INSERT INTO `users`(.*) VALUES (.*)"
			mock.ExpectExec(insertQuery).WillReturnResult(
				sqlmock.NewResult(1, 1),
			)

			query := regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? LIMIT 1")
			mock.ExpectQuery(query).WillReturnRows(
				sqlmock.NewRows([]string{"ID", "Username", "Password"}).
					AddRow(
						"01ENSP2J4H9QRMWYJ3ZQTYEGYJ",
						"username",
						"$argon2id$v=19$m=65536,t=8,p=8$rWzlADlkY3v8OsLwVKYFGA$owGB0DwHIFywicD4Ydqz2FY5LT406eNUD7aCgMytqZlmuWnHWkNK3Nl5R+aAhV5tVWOqUObqZaO5+zsGwMNBAA",
					),
			)
			runFunc(t,
				fields{baseRepository{}},
				args{
					ctx: ctx,
					ent: entity.User{
						Username: "username",
						Password: "$argon2id$v=19$m=65536,t=8,p=8$rWzlADlkY3v8OsLwVKYFGA$owGB0DwHIFywicD4Ydqz2FY5LT406eNUD7aCgMytqZlmuWnHWkNK3Nl5R+aAhV5tVWOqUObqZaO5+zsGwMNBAA",
					},
				},
				wants{
					wantCreated: entity.User{},
					wantErr:     true,
				},
			)
		})
	})
}
