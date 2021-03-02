package main

import (
	"os"
	"os/signal"

	"github.com/lovung/GoCleanArchitecture/app/config"
	"github.com/lovung/GoCleanArchitecture/app/external/api"
	"github.com/lovung/GoCleanArchitecture/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

type application struct {
	cfg    *config.Config
	engine *gin.Engine
}

func (s *application) start() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	s.initDBConnection(s.cfg.MySQL)
	s.initJWTSession(s.cfg.JWTSecret)

	s.engine = api.Restful(s.cfg)

	go func() {
		if err := s.engine.Run(":" + s.cfg.HTTPServer.Port); err != nil {
			panic(err)
		}
	}()

	<-interrupt

	s.stopping()

	return nil
}

// stopping will stop running job or release resources the server was used
func (s *application) stopping() {
	logger.Debug("server stopped")
}

func newService(ctx *cli.Context) *application {
	s := &application{}
	s.loadConfig(ctx)

	logger.Init(s.cfg.LogLevel == string(logger.DebugLevel))
	logger.SetLevel(s.cfg.LogLevel)
	return s
}

func (s *application) loadConfig(ctx *cli.Context) {
	conf := &config.Config{
		Env: ctx.String(EnvFlag.Name),
		HTTPServer: config.ServerCfg{
			Port:    ctx.String(HTTPPortFlag.Name),
			Timeout: ctx.Duration(HTTPTimeoutFlag.Name),
		},
		MySQL: config.MySQL{
			ConnectionString: ctx.String(MYSQLConnFlag.Name),
			Host:             ctx.String(MYSQLHostFlag.Name),
			Port:             ctx.String(MySQLPortFlag.Name),
			User:             ctx.String(MySQLUserFlag.Name),
			Password:         ctx.String(MySQLPasswordFlag.Name),
			DB:               ctx.String(MySQLDatabaseFlag.Name),
			MaxOpenConns:     ctx.Int(MySQLMaxOpenConnsFlag.Name),
			MaxIdleConns:     ctx.Int(MySQLMaxIdleConnsFlag.Name),
			ConnMaxLifetime:  ctx.Int(MySQLConnMaxLifetimeFlag.Name),
			IsEnabledLog:     ctx.String(LogLevelFlag.Name) == string(logger.DebugLevel),
		},
		LogLevel:         ctx.String(LogLevelFlag.Name),
		JWTSecret:        ctx.String(JWTSecretFlag.Name),
		EnabledProfiling: ctx.Bool(EnabledProfilingFlag.Name),
	}

	s.cfg = conf
	config.SetConfig(conf)
}
