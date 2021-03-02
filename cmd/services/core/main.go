package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	flags := []cli.Flag{
		EnvFlag,
		AppNameFlag,
		AppVersionFlag,
		MYSQLConnFlag,
		MYSQLHostFlag,
		MySQLPortFlag,
		MySQLUserFlag,
		MySQLPasswordFlag,
		MySQLDatabaseFlag,
		MySQLMaxOpenConnsFlag,
		MySQLMaxIdleConnsFlag,
		MySQLConnMaxLifetimeFlag,
		HTTPPortFlag,
		HTTPTimeoutFlag,
		LogLevelFlag,
		JWTSecretFlag,
		EnabledProfilingFlag,
	}

	app := &cli.App{
		Name:  "github.com/lovung/GoCleanArchitecture core Service",
		Flags: flags,
		Action: func(ctx *cli.Context) error {
			srv := newService(ctx)

			if err := srv.start(); err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
