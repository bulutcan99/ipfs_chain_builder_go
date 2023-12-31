package config_builder

import (
	"fmt"
	"github.com/bulutcan99/go_ipfs_chain_builder/pkg/env"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB_HOST     = &env.Env.DbHost
	DB_PORT     = &env.Env.DbPort
	DB_USER     = &env.Env.DbUser
	DB_PASSWORD = &env.Env.DbPassword
	DB_NAME     = &env.Env.DbName
)

func ConnectionURLBuilder(n string) (string, error) {
	var url string
	switch n {
	case "mysql":
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			*DB_USER,
			*DB_PASSWORD,
			*DB_HOST,
			*DB_PORT,
			*DB_NAME,
		)
		fmt.Println("URL:", url)

	default:
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	return url, nil
}
