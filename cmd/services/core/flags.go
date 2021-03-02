package main

import (
	"time"

	"github.com/urfave/cli/v2"
)

// APP env: should only put environment variable related to the service itself, eg: Application Name, version, running environment, ...
var (
	EnvFlag = &cli.StringFlag{
		Name:    "env",
		Usage:   "Application environment: development, staging, production",
		EnvVars: []string{"ENV"},
		Value:   "development",
	}

	AppNameFlag = &cli.StringFlag{
		Name:    "app_name",
		Usage:   "Application name",
		EnvVars: []string{"APP_NAME"},
		Value:   "GoCleanArchitecture core service",
	}

	AppVersionFlag = &cli.StringFlag{
		Name:    "app_version",
		Usage:   "Application version",
		EnvVars: []string{"APP_VERSION"},
		Value:   "v1",
	}

	HTTPPortFlag = &cli.StringFlag{
		Name:    "http_port",
		Usage:   "Port binding to application",
		EnvVars: []string{"HTTP_PORT"},
		Value:   "20000",
	}

	HTTPTimeoutFlag = &cli.DurationFlag{
		Name:    "http_timeout",
		Usage:   "Time out for HTTP request",
		EnvVars: []string{"HTTP_TIMEOUT"},
		Value:   time.Second,
	}
)

// MYSQL env
var (
	MYSQLConnFlag = &cli.StringFlag{
		Name:    "mysql_conn",
		Usage:   `specify MySQL connection string. If non empty then other flags begin with "mysql_" will be ignore`,
		EnvVars: []string{"JAWSDB_URL"}, // support for heroku deployment.
		Value:   "",
	}

	MYSQLHostFlag = &cli.StringFlag{
		Name:    "mysql_host",
		Usage:   "specify MySQL host",
		EnvVars: []string{"MYSQL_HOST"},
		Value:   "localhost",
	}

	MySQLPortFlag = &cli.StringFlag{
		Name:    "mysql_port",
		Usage:   "MySQL port is using by application",
		EnvVars: []string{"MYSQL_PORT"},
		Value:   "10001",
	}

	MySQLUserFlag = &cli.StringFlag{
		Name:    "mysql_user",
		Usage:   "specify MySQL user",
		EnvVars: []string{"MYSQL_USER"},
		Value:   "admin",
	}

	MySQLPasswordFlag = &cli.StringFlag{
		Name:    "mysql_password",
		Usage:   "password used for MySQL user",
		EnvVars: []string{"MYSQL_PASSWORD"},
		Value:   "gocleanarchitecture12345",
	}

	MySQLDatabaseFlag = &cli.StringFlag{
		Name:    "mysql_database",
		Usage:   "MySQL database is using by application",
		EnvVars: []string{"MYSQL_DATABASE"},
		Value:   "gocleanarchitecture",
	}

	MySQLMaxOpenConnsFlag = &cli.IntFlag{
		Name:    "mysql_max_open_conns",
		Usage:   "sets the maximum number of open connections to the database",
		EnvVars: []string{"MYSQL_MAX_OPEN_CONNS"},
		Value:   10,
	}

	MySQLMaxIdleConnsFlag = &cli.IntFlag{
		Name:    "mysql_max_idle_conns",
		Usage:   "sets the maximum number of connections in the idle connection pool",
		EnvVars: []string{"MYSQL_MAX_IDLE_CONNS"},
		Value:   5,
	}

	MySQLConnMaxLifetimeFlag = &cli.IntFlag{
		Name:    "mysql_conn_max_lifetime",
		Usage:   "sets the maximum amount of time in minutes a connection may be reused",
		EnvVars: []string{"MYSQL_CONN_MAX_LIFETIME"},
		Value:   60,
	}
)

// // Local storage
// var (
// 	LocalDataDirFlag = &cli.StringFlag{
// 		Name:    "local_data_dir",
// 		Usage:   "local data directory path",
// 		EnvVars: []string{"LOCAL_DATA_DIR"},
// 		Value:   "github.com/lovung/GoCleanArchitecture",
// 	}
// )

// // Redis Config flag
// var (
// 	RedisConnFlag = &cli.StringFlag{
// 		Name:    "redis_conn",
// 		Usage:   `specify Redis connection string. If non empty then other flags begin with "redis_" will be ignore`,
// 		EnvVars: []string{"REDISCLOUD_URL"}, // support for heroku deployment.
// 		Value:   "",
// 	}

// 	RedisHostFlag = &cli.StringFlag{
// 		Name:    "redis_host",
// 		Usage:   "specify Redis host",
// 		EnvVars: []string{"REDIS_HOST"},
// 		Value:   "localhost",
// 	}

// 	RedisPortFlag = &cli.StringFlag{
// 		Name:    "redis_port",
// 		Usage:   "Redis port is using by application",
// 		EnvVars: []string{"REDIS_PORT"},
// 		Value:   "16379",
// 	}

// 	RedisUserFlag = &cli.StringFlag{
// 		Name:    "redis_user",
// 		Usage:   "specify Redis user",
// 		EnvVars: []string{"REDIS_USER"},
// 		Value:   "admin",
// 	}

// 	RedisPasswordFlag = &cli.StringFlag{
// 		Name:    "redis_password",
// 		Usage:   "password used for Redis user",
// 		EnvVars: []string{"REDIS_PASSWORD"},
// 		Value:   "moneyforward123",
// 	}

// 	RedisEnabledTLSFlag = &cli.BoolFlag{
// 		Name:    "redis_enabled_tls",
// 		Usage:   "enable tls for Redis tls connection",
// 		EnvVars: []string{"REDIS_ENABLED_TLS"},
// 		Value:   false,
// 	}

// 	RedisInsecureSkipVerifyFlag = &cli.BoolFlag{
// 		Name:    "redis_insecure_skip_verify",
// 		Usage:   "insecure_skip_verify used for Redis tls verify",
// 		EnvVars: []string{"REDIS_INSECURE_SKIP_VERIFY"},
// 		Value:   true,
// 	}

// 	RedisDatabaseFlag = &cli.IntFlag{
// 		Name:    "redis_db",
// 		Usage:   "Redis database is using by application",
// 		EnvVars: []string{"REDIS_DB"},
// 		Value:   0,
// 	}

// 	RedisPoolSizeFlag = &cli.IntFlag{
// 		Name:    "redis_max_open_conns",
// 		Usage:   "sets the maximum number of open connections to the database",
// 		EnvVars: []string{"REDIS_POOL_SIZE"},
// 		Value:   10,
// 	}
// )

// Log and notifier env
var (
	LogLevelFlag = &cli.StringFlag{
		Name:    "log_level",
		Usage:   "Level to log message to standard loggger: panic, fatal, error, warn, warning, info, debug",
		EnvVars: []string{"LOG_LEVEL"},
		Value:   "info",
	}
)

// For Authentication of JWT
var (
	JWTSecretFlag = &cli.StringFlag{
		Name:    "jwt_secret",
		Usage:   "it will be used to sign JWT",
		EnvVars: []string{"JWT_SECRET"},
		Value:   "dummy_for_local",
	}
)

// For pprof middleware
var (
	EnabledProfilingFlag = &cli.BoolFlag{
		Name:    "enabled_pprof",
		Usage:   "enable pprof middleware",
		EnvVars: []string{"ENABLED_PPROF"},
		Value:   false,
	}
)
