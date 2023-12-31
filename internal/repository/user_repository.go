package repository

import (
	"context"
	"github.com/bulutcan99/go_ipfs_chain_builder/model"
	config_mysql "github.com/bulutcan99/go_ipfs_chain_builder/pkg/config/mysql"
	"github.com/jmoiron/sqlx"
)

type IUserRepo interface {
	AddUser(user model.User) (int64, error)
	GetUser(id int) (model.User, error)
}

type UserRepo struct {
	db      *sqlx.DB
	context context.Context
}

func NewUserRepo(db *config_mysql.MYSQL) *UserRepo {
	return &UserRepo{
		db:      db.Client,
		context: db.Context,
	}
}

func (r *UserRepo) AddUser(user model.User) (int64, error) {
	result, err := r.db.NamedExec("INSERT INTO users (name) VALUES (:name)", user)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) GetUser(id int) (model.User, error) {
	var user model.User

	err := r.db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return user, err
	}

	return user, nil
}
