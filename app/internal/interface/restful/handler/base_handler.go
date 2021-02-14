package handler

import (
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/presenter"

	"github.com/gin-gonic/gin"
)

// BaseHandler help us respond to client
type BaseHandler struct{}

// SetMeta to put meta information into context
func (h *BaseHandler) SetMeta(ctx *gin.Context, meta presenter.MetaResponse) {
	ctx.Set(presenter.MetaContextKey.String(), meta)
}

// SetData to put data information into context
func (h *BaseHandler) SetData(ctx *gin.Context, data interface{}) {
	ctx.Set(presenter.DataContextKey.String(), data)
}

// SetError to put meta information into context
func (h *BaseHandler) SetError(ctx *gin.Context, err error) {
	ctx.Set(presenter.ErrorContextKey.String(), err)
}
