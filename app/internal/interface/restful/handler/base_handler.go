package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lovung/GoCleanArchitecture/app/internal/appctx"
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/presenter"
)

// BaseHandler help us respond to client
type BaseHandler struct{}

// SetMeta to put meta information into context
func (h *BaseHandler) SetMeta(ctx *gin.Context, meta presenter.MetaResponse) {
	newCtx := appctx.SetValue(ctx.Request.Context(), appctx.MetaContextKey, meta)
	ctx.Request = ctx.Request.WithContext(newCtx)
}

// SetData to put data information into context
func (h *BaseHandler) SetData(ctx *gin.Context, data interface{}) {
	newCtx := appctx.SetValue(ctx.Request.Context(), appctx.DataContextKey, data)
	ctx.Request = ctx.Request.WithContext(newCtx)
}

// SetError to put meta information into context
func (h *BaseHandler) SetError(ctx *gin.Context, err error) {
	newCtx := appctx.SetValue(ctx.Request.Context(), appctx.ErrorContextKey, err)
	ctx.Request = ctx.Request.WithContext(newCtx)
}
