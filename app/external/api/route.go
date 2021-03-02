package api

import (
	"net/http"
	"time"

	"github.com/lovung/GoCleanArchitecture/app/config"
	"github.com/lovung/GoCleanArchitecture/app/internal/interface/restful/middleware"
	"github.com/lovung/GoCleanArchitecture/app/registry"
	"github.com/lovung/GoCleanArchitecture/pkg/logger"

	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

const (
	envStaging    = "staging"
	envProduction = "production"
)

//nolint:funlen
// Restful define mapping routes
// @title github.com/lovung/GoCleanArchitecture core service
// @version 1.0
// @description This is the project of github.com/lovung/GoCleanArchitecture
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func Restful(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(logger.Instance(), time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(logger.Instance(), true))
	if cfg.Env != envProduction && cfg.Env != envStaging {
		router.Use(middleware.CorsMiddleware())
	}
	router.Use(middleware.AddTimeout)
	router.GET("/", root)
	router.GET("/api/healthz", healthz)
	router.Use(middleware.JSONWriterMiddleware)
	// pprof: default path is /debug/pprof
	if cfg.EnabledProfiling {
		pprof.Register(router)
	}

	authHandler := registry.AuthHandler()
	txnMw := registry.TransactionMiddleware()

	authRoute := router.Group("/auth")
	authRoute.Use(txnMw.StartRequest)
	authRoute.Use(txnMw.EndRequest)
	authRoute.POST("/register", authHandler.Register)

	return router
}

func root(ctx *gin.Context) {
	type svcInfo struct {
		JSONAPI struct {
			Version string `json:"version,omitempty"`
			Name    string `json:"name,omitempty"`
		} `json:"jsonapi"`
	}

	info := svcInfo{}
	info.JSONAPI.Version = "v1"
	info.JSONAPI.Name = "github.com/lovung/GoCleanArchitecture core API"

	ctx.JSON(http.StatusOK, info)
}

func healthz(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}
