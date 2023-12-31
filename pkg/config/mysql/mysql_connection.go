package config_psql

import (
	"context"
	"github.com/bulutcan99/go_ipfs_chain_builder/pkg/config"
	"github.com/bulutcan99/go_ipfs_chain_builder/pkg/env"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"sync"
	"time"
)

var (
	doOnce               sync.Once
	client               *sqlx.DB
	DB_MAX_CON           = &env.Env.DbMaxConnections
	DB_MAX_IDLE_CON      = &env.Env.DbMaxIdleConnections
	DB_MAX_LIFE_TIME_CON = &env.Env.DbMaxLifetimeConnections
)

type MYSQL struct {
	Client  *sqlx.DB
	Context context.Context
}

func NewMYSQLConnection() *MYSQL {
	ctx := context.Background()
	postgresConnURL, err := config_builder.ConnectionURLBuilder("mysql")
	if err != nil {
		panic(err)
	}
	doOnce.Do(func() {
		db, err := sqlx.ConnectContext(ctx, "mysql", postgresConnURL)
		if err != nil {
			panic(err)
		}

		db.SetMaxOpenConns(*DB_MAX_CON)
		db.SetMaxIdleConns(*DB_MAX_IDLE_CON)
		db.SetConnMaxLifetime(time.Duration(*DB_MAX_LIFE_TIME_CON))
		if err := db.PingContext(ctx); err != nil {
			panic(err)
		}
		client = db
	})

	zap.S().Infof("Connected to Postgres successfully.")
	return &MYSQL{
		Client:  client,
		Context: ctx,
	}
}

func (pg *MYSQL) Close() {
	err := pg.Client.Close()
	if err != nil {
		zap.S().Errorf("Error while closing the database connection: %s", err)
	}

	zap.S().Infof("Connection to Postgres closed successfully")
}
