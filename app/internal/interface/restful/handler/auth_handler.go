package handler

import (
	"net/http"

	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/presenter"
	"github.com/lovung/GoCleanArchitecture/app/internal/usecase"
	"github.com/lovung/GoCleanArchitecture/app/internal/usecase/dto"
	"github.com/lovung/GoCleanArchitecture/pkg/copier"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles the authentication/authorization requests
type AuthHandler struct {
	BaseHandler
	userUseCase usecase.UserUseCase
}

// NewAuthHandler constructor
func NewAuthHandler(
	userUseCase usecase.UserUseCase,
) AuthHandler {
	return AuthHandler{
		userUseCase: userUseCase,
	}
}

// Register function in handler to register API
// @Summary Register the new user
// @Description Register the new user
// @Tags auth
// @Accept   json
// @Produce  json
// @Param 	 Authorization	header	string true "Bearer jwt_access_token"
// @Param	 body   body   presenter.RegisterRequest true    "Body of request"
// @Success 201 {object} presenter.UserInformation
// @Failure 400 {object} presenter.Error
// @Router /register [post]
func (hdl *AuthHandler) Register(ctx *gin.Context) {
	var (
		req    presenter.RegisterRequest
		res    presenter.UserInformation
		dtoReq dto.CreateUserRequest
		err    error
	)
	defer func() {
		hdl.SetError(ctx, err)
	}()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	copier.MustCopy(&dtoReq, &req)
	resDto, err := hdl.userUseCase.Create(ctx.Request.Context(), dtoReq)
	if err != nil {
		return
	}
	copier.MustCopy(&res, &resDto)
	hdl.SetData(ctx, res)
	hdl.SetMeta(ctx, presenter.MetaResponse{
		Code: http.StatusCreated,
	})
}
