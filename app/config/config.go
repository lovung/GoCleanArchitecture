package config

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var (
	once      sync.Once
	singleton *Config
)

// SetConfig to set configuration of service.
func SetConfig(cfg *Config) *Config {
	once.Do(func() {
		singleton = cfg
	})
	return singleton
}

// GetConfig gets the instance of singleton
func GetConfig() *Config {
	return singleton
}

// MySQL config for MySQL server
type MySQL struct {
	ConnectionString string // if this field is non empty then other fields will be ignore when building connection string
	Host             string
	User             string
	Password         string
	DB               string
	Port             string
	MaxOpenConns     int
	MaxIdleConns     int
	ConnMaxLifetime  int // time in minute
	IsEnabledLog     bool
}

// Conn return connection string
func (m *MySQL) Conn() string {
	if m.ConnectionString != "" {
		connStr := strings.TrimPrefix(m.ConnectionString, "mysql://")
		connStr = strings.Split(connStr, "?")[0]
		connStr = strings.Replace(connStr, "@", "@tcp(", 1)
		if strings.Contains(connStr, ":3306") {
			connStr = strings.Replace(connStr, "/", ")/", 1)
		} else {
			connStr = strings.Replace(connStr, "/", ":3306)/", 1)
		}
		connStr += "?parseTime=true&charset=utf8mb4"

		return connStr
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		m.User, m.Password, m.Host, m.Port, m.DB)
}

// ServerCfg server addresses
type ServerCfg struct {
	Port    string
	Timeout time.Duration
}

// Config  is APP config information
type Config struct {
	Env              string
	HTTPServer       ServerCfg
	MySQL            MySQL
	LogLevel         string
	JWTSecret        string
	EnabledProfiling bool
}
