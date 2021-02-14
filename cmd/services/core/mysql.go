package main

import (
	"os"
	"time"

	"github.com/lovung/GoCleanArchitecture/app/config"
	"github.com/lovung/GoCleanArchitecture/pkg/gormutil"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (s *application) initDBConnection(cfg config.MySQL) *gorm.DB {
	var db *gorm.DB
	logMode := logger.Default.LogMode(logger.Error)
	if cfg.IsEnabledLog {
		logMode = logger.Default.LogMode(logger.Info)
	}
	db, err := gormutil.OpenDBConnection(cfg.Conn(),
		gorm.Config{
			Logger: logMode,
		},
	)
	if err != nil {
		os.Exit(1)
	}
	gormer, err := db.DB()
	if err != nil {
		os.Exit(1)
	}
	gormer.SetMaxOpenConns(cfg.MaxOpenConns)
	gormer.SetMaxIdleConns(cfg.MaxIdleConns)
	gormer.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)
	return db
}
