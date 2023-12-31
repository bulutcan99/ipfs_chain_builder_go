package service

import (
	"github.com/bulutcan99/go_ipfs_chain_builder/internal/repository"
	"github.com/bulutcan99/go_ipfs_chain_builder/model"
)

type UserService struct {
	userService repository.IUserRepo
}

func NewUserService(userService repository.IUserRepo) *UserService {
	return &UserService{
		userService: userService,
	}
}

func (us *UserService) AddUser(user model.User) (int64, error) {
	return us.userService.AddUser(user)
}

func (us *UserService) GetUser(id int) (model.User, error) {
	return us.userService.GetUser(id)
}
