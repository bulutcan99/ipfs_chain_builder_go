package env

import (
	"fmt"
	custom_error "github.com/bulutcan99/go_ipfs_chain_builder/pkg/error"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type ENV struct {
	DbHost                   string `env:"DB_HOST,required"`
	DbPort                   int    `env:"DB_PORT,required"`
	DbUser                   string `env:"DB_USER,required"`
	DbPassword               string `env:"DB_PASSWORD,required"`
	DbName                   string `env:"DB_NAME,required"`
	DbTable                  string `env:"DB_TABLE,required"`
	DbMaxConnections         int    `env:"DB_MAX_CONNECTIONS,required"`
	DbMaxIdleConnections     int    `env:"DB_MAX_IDLE_CONNECTIONS,required"`
	DbMaxLifetimeConnections int    `env:"DB_MAX_LIFETIME_CONNECTIONS,required"`
	IpfsPort                 int    `env:"IPFS_PORT,required"`
	LogLevel                 string `env:"LOG_LEVEL,required"`
}

var doOnce sync.Once
var Env ENV

func ParseEnv() *ENV {
	doOnce.Do(func() {
		e := godotenv.Load()
		if e != nil {
			fmt.Println(e)
			err := custom_error.ParseError()
			if err != nil {
				fmt.Println(err.Error())
			}
			os.Exit(1)
		}
		if err := env.Parse(&Env); err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(0)
		}
	})
	return &Env
}
