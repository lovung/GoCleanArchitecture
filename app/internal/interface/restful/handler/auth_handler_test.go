package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/middleware"
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/presenter"
	"github.com/lovung/GoCleanArchitecture/app/internal/usecase"
	"github.com/lovung/GoCleanArchitecture/app/internal/usecase/dto"
	"github.com/lovung/GoCleanArchitecture/app/internal/usecase/mockusecase"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthHandler(t *testing.T) {
	type args struct {
		userUseCase usecase.UserUseCase
	}
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mUserUseCase := mockusecase.NewMockUserUseCase(mockCtrl)

	testCases := []struct {
		name string
		args args
		want AuthHandler
	}{
		{
			args: args{
				userUseCase: mUserUseCase,
			},
			want: AuthHandler{
				BaseHandler: BaseHandler{},
				userUseCase: mUserUseCase,
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthHandler(tt.args.userUseCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthHandler_Register(t *testing.T) {
	type fields struct {
		BaseHandler BaseHandler
		userUseCase usecase.UserUseCase
	}
	type args struct {
		// ctx *gin.Context
		req presenter.RegisterRequest
	}
	type wants struct {
		httpCode int
		body     interface{}
	}
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mUserUseCase := mockusecase.NewMockUserUseCase(mockCtrl)
	_fields := fields{
		BaseHandler: BaseHandler{},
		userUseCase: mUserUseCase,
	}

	runFunc := func(t *testing.T, fields fields, args args, w wants) {
		hdl := &AuthHandler{
			BaseHandler: fields.BaseHandler,
			userUseCase: fields.userUseCase,
		}
		// Switch to test mode so you don't get such noisy output
		gin.SetMode(gin.TestMode)

		// Create a response recorder so you can inspect the response
		resp := httptest.NewRecorder()
		r := gin.Default()
		// Setup your router, just like you did in your main function, and
		// register your routes
		r.Use(middleware.JSONWriterMiddleware)
		r.POST("/register", hdl.Register)
		// Create the mock request you'd like to test. Make sure the second argument
		// here is the same as one of the routes you defined in the router setup
		// block!
		body, err := json.Marshal(args.req)
		assert.NoError(t, err)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
		assert.NoError(t, err)

		// Perform the request
		r.ServeHTTP(resp, req)
		assert.Equal(t, w.httpCode, resp.Code)
		if diff := cmp.Diff(w.body, resp.Body.String()); diff != "" {
			t.Errorf("Diff body = %v", diff)
			assert.Equal(t, w.body, resp.Body.String())
		}
	}
	t.Run("#1: Error when binding", func(t *testing.T) {
		runFunc(t, _fields,
			args{
				req: presenter.RegisterRequest{
					Username: "username",
				},
			},
			wants{
				httpCode: 400,
				body:     "{\"meta\":{\"code\":400,\"message\":\"error when binding the request\"},\"errors\":[{\"code\":422,\"detail\":\"Key: 'RegisterRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag\",\"source\":{\"pointer\":\"RegisterRequest.Password\",\"parameter\":\"Password\"}}]}",
			})
	})
	t.Run("#2: Error when binding username", func(t *testing.T) {
		runFunc(t, _fields,
			args{
				req: presenter.RegisterRequest{
					Password: "password",
				},
			},
			wants{
				httpCode: 400,
				body:     "{\"meta\":{\"code\":400,\"message\":\"error when binding the request\"},\"errors\":[{\"code\":422,\"detail\":\"Key: 'RegisterRequest.Username' Error:Field validation for 'Username' failed on the 'required' tag\",\"source\":{\"pointer\":\"RegisterRequest.Username\",\"parameter\":\"Username\"}}]}",
			})
	})
	t.Run("#3: Call use case got error", func(t *testing.T) {
		mUserUseCase.EXPECT().Create(gomock.Any(), dto.CreateUserRequest{
			Password: "password",
			Username: "username",
		}).Return(dto.OneUserResponse{}, errors.New("error"))
		runFunc(t, _fields,
			args{
				req: presenter.RegisterRequest{
					Username: "username",
					Password: "password",
				},
			},
			// TODO: Update wants when implementing the JSONWriterMiddleware
			// to handle the App Error (error from business logic)
			wants{
				httpCode: 204,
				body:     "",
			})
	})
	t.Run("#4: Success", func(t *testing.T) {
		mUserUseCase.EXPECT().Create(gomock.Any(), dto.CreateUserRequest{
			Password: "password",
			Username: "username",
		}).Return(dto.OneUserResponse{
			ID:       "id",
			Username: "username",
		}, nil)
		runFunc(t, _fields,
			args{
				req: presenter.RegisterRequest{
					Username: "username",
					Password: "password",
				},
			},
			wants{
				httpCode: 201,
				body:     "{\"meta\":{\"code\":201},\"data\":{\"id\":\"id\",\"username\":\"username\"}}",
			})
	})
}

// func TestAuthHandler_RegisterWithTransaction(t *testing.T) {
// 	type fields struct {
// 		BaseHandler BaseHandler
// 		userUseCase usecase.UserUseCase
// 	}
// 	type args struct {
// 		// ctx *gin.Context
// 		req presenter.RegisterRequest
// 	}
// 	type wants struct {
// 		httpCode int
// 		body     interface{}
// 	}
// 	t.Parallel()
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()
// 	mUserUseCase := mockusecase.NewMockUserUseCase(mockCtrl)
// 	mTransactionManager := mocktrans.NewMockManager(mockCtrl)
// 	gDB, _, err := tests.OpenDBConnection()
// 	logger.Init(true)
// 	assert.NoError(t, err)
// 	_fields := fields{
// 		BaseHandler: BaseHandler{},
// 		userUseCase: mUserUseCase,
// 	}

// 	runFunc := func(t *testing.T, fields fields, args args, w wants) {
// 		hdl := &AuthHandler{
// 			BaseHandler: fields.BaseHandler,
// 			userUseCase: fields.userUseCase,
// 		}
// 		txnMw := middleware.NewTransactionMiddleware(mTransactionManager)
// 		// Switch to test mode so you don't get such noisy output
// 		gin.SetMode(gin.TestMode)

// 		// Create a response recorder so you can inspect the response
// 		resp := httptest.NewRecorder()
// 		ctx, r := gin.CreateTestContext(resp)
// 		// Setup your router, just like you did in your main function, and
// 		// register your routes
// 		r.Use(middleware.JSONWriterMiddleware)
// 		r.Use(txnMw.StartRequest)
// 		r.Use(txnMw.EndRequest)
// 		r.POST("/register", hdl.Register)
// 		// Create the mock request you'd like to test. Make sure the second argument
// 		// here is the same as one of the routes you defined in the router setup
// 		// block!
// 		body, err := json.Marshal(args.req)
// 		assert.NoError(t, err)
// 		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
// 		assert.NoError(t, err)

// 		tx := gDB.Begin()
// 		newCtx := appctx.SetValue(ctx.Request.Context(), appctx.TransactionContextKey, tx)

// 		mTransactionManager.EXPECT().TxnBegin(gomock.Any()).Return(newCtx)
// 		mTransactionManager.EXPECT().TxnRollback(gomock.Any()).Return(nil)
// 		// Perform the request
// 		r.ServeHTTP(resp, req)
// 		assert.Equal(t, w.httpCode, resp.Code)
// 		if diff := cmp.Diff(w.body, resp.Body.String()); diff != "" {
// 			t.Errorf("Diff body = %v", diff)
// 			assert.Equal(t, w.body, resp.Body.String())
// 		}
// 	}
// 	t.Run("#1: Error when binding", func(t *testing.T) {
// 		runFunc(t, _fields,
// 			args{
// 				req: presenter.RegisterRequest{
// 					Username: "username",
// 				},
// 			},
// 			wants{
// 				httpCode: 400,
// 				body:     "{\"meta\":{\"code\":400,\"message\":\"error when binding the request\"},\"errors\":[{\"code\":422,\"detail\":\"Key: 'RegisterRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag\",\"source\":{\"pointer\":\"RegisterRequest.Password\",\"parameter\":\"Password\"}}]}",
// 			})
// 	})
// t.Run("#2: Error when binding username", func(t *testing.T) {
// 	ctx := context.Background()
// 	tx := gDB.Begin()
// 	ctx = appctx.SetValue(ctx, appctx.TransactionContextKey, tx)
// 	mTransactionManager.EXPECT().TxnBegin(gomock.Any()).Return(ctx)
// 	mTransactionManager.EXPECT().TxnRollback(gomock.Any()).Return(nil)
// 	runFunc(t, _fields,
// 		args{
// 			req: presenter.RegisterRequest{
// 				Password: "password",
// 			},
// 		},
// 		wants{
// 			httpCode: 400,
// 			body:     "{\"meta\":{\"code\":400,\"message\":\"error when binding the request\"},\"errors\":[{\"code\":422,\"detail\":\"Key: 'RegisterRequest.Username' Error:Field validation for 'Username' failed on the 'required' tag\",\"source\":{\"pointer\":\"RegisterRequest.Username\",\"parameter\":\"Username\"}}]}",
// 		})
// })
// t.Run("#3: Call use case got error", func(t *testing.T) {
// 	ctx := context.Background()
// 	tx := gDB.Begin()
// 	ctx = appctx.SetValue(ctx, appctx.TransactionContextKey, tx)
// 	mTransactionManager.EXPECT().TxnBegin(gomock.Any()).Return(ctx)
// 	mTransactionManager.EXPECT().TxnRollback(gomock.Any()).Return(nil)
// 	mUserUseCase.EXPECT().Create(gomock.Any(), dto.CreateUserRequest{
// 		Password: "password",
// 		Username: "username",
// 	}).Return(dto.OneUserResponse{}, errors.New("error"))
// 	runFunc(t, _fields,
// 		args{
// 			req: presenter.RegisterRequest{
// 				Username: "username",
// 				Password: "password",
// 			},
// 		},
// 		// TODO: Update wants when implementing the JSONWriterMiddleware
// 		// to handle the App Error (error from business logic)
// 		wants{
// 			httpCode: 204,
// 			body:     "",
// 		})
// })
// t.Run("#4: Success", func(t *testing.T) {
// 	ctx := context.Background()
// 	tx := gDB.Begin()
// 	ctx = appctx.SetValue(ctx, appctx.TransactionContextKey, tx)
// 	mTransactionManager.EXPECT().TxnBegin(gomock.Any()).Return(ctx)
// 	mTransactionManager.EXPECT().TxnCommit(gomock.Any()).Return(nil)
// 	mTransactionManager.EXPECT().GetTxn(gomock.Any()).Return(tx)
// 	mUserUseCase.EXPECT().Create(gomock.Any(), dto.CreateUserRequest{
// 		Password: "password",
// 		Username: "username",
// 	}).Return(dto.OneUserResponse{
// 		ID:       "id",
// 		Username: "username",
// 	}, nil)
// 	runFunc(t, _fields,
// 		args{
// 			req: presenter.RegisterRequest{
// 				Username: "username",
// 				Password: "password",
// 			},
// 		},
// 		wants{
// 			httpCode: 201,
// 			body:     "{\"meta\":{\"code\":201},\"data\":{\"id\":\"id\",\"username\":\"username\"}}",
// 		})
// })
// }
