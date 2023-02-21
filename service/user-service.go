package service

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type UserService interface {
	FindByPassEmail(email, password string) (entity.AuthUser, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService() UserService {
	repo := repository.NewUserRepository()

	return &userService{
		userRepo: repo,
	}
}

func (userServ *userService) FindByPassEmail(email, password string) (entity.AuthUser, error) {
	return userServ.userRepo.FindByPassEmail(email, password)
}
