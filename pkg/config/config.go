package config

import (
	"sync"

	"github.com/lovung/GoCleanArchitecture/app/config"
)

var (
	once      sync.Once
	singleton *config.Config
)

// SetConfig to set configuration of service.
func SetConfig(cfg *config.Config) *config.Config {
	once.Do(func() {
		singleton = cfg
	})
	return singleton
}

// GetConfig gets the instance of singleton
func GetConfig() *config.Config {
	return singleton
}
