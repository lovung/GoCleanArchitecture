package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/lovung/GoCleanArchitecture/app/config"
)

// AddTimeout for context
func AddTimeout(ctx *gin.Context) {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctxWTimeout, cancel := context.WithTimeout(
		ctx.Request.Context(),
		config.GetConfig().HTTPServer.Timeout,
	)
	ctx.Request = ctx.Request.WithContext(ctxWTimeout)
	defer cancel()
	ctx.Next()
}
