package middleware

import (
	"net/http"

	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/presenter"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// JSONWriterMiddleware follow jsonapi.org
// Respond check errors, prepare meta and respond data
func JSONWriterMiddleware(ctx *gin.Context) {
	ctx.Next()

	// Check error if exists
	// Base on error/success to return meta object
	var (
		res      presenter.Response
		httpCode int
	)
	appErr, ok := ctx.Get(presenter.ErrorContextKey.String())
	if ok && appErr != nil {
		_processAppError(&res, appErr)
		httpCode = res.Meta.Code
	}

	// Respond the data object/array
	data, ok := ctx.Get(presenter.DataContextKey.String())
	if ok {
		res.Data = data
	}
	meta, ok := ctx.Get(presenter.MetaContextKey.String())
	if ok && meta != nil {
		metaRes, ok1 := meta.(presenter.MetaResponse)
		if ok1 {
			res.Meta = metaRes
			httpCode = metaRes.Code
		}
	}
	if res.IsEmpty() {
		ctx.JSON(http.StatusNoContent, nil)
	} else {
		ctx.JSON(httpCode, res)
	}
}

func _processAppError(res *presenter.Response, appErr interface{}) {
	if appErr == nil {
		return
	}
	bindingErr := _catchBindingError(appErr.(error))
	if bindingErr != nil {
		res.Errors = bindingErr.(presenter.ErrorResponses)
		res.Meta = presenter.MetaResponse{
			Code:    http.StatusBadRequest,
			Message: "error when binding the request",
		}
	}
}

func _catchBindingError(appErr error) error {
	if appErr == nil {
		return nil
	}
	var errs presenter.ErrorResponses
	if _, ok := appErr.(*validator.InvalidValidationError); ok {
		errs.Append(presenter.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Detail: "invalid validation error",
		})
		return errs
	}

	if vldrErr, ok := appErr.(validator.ValidationErrors); ok {
		errs.FromValidationErrors(vldrErr)
		return errs
	}

	return nil
}
