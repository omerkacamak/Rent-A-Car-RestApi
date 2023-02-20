package service

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type UserService interface {
	Save(user entity.AuthUser) error
	Update(user entity.AuthUser) error
	Delete(user entity.AuthUser) error
	FindAll() []entity.AuthUser

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

func (userServ *userService) Save(user entity.AuthUser) error {
	err := userServ.userRepo.Save(user)
	if err != nil {
		return err
	}
	return nil
}
func (userServ *userService) Update(user entity.AuthUser) error {
	err := userServ.userRepo.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (userServ *userService) Delete(user entity.AuthUser) error {
	err := userServ.userRepo.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
func (userServ *userService) FindAll() []entity.AuthUser {
	result := userServ.userRepo.FindAll()

	return result
}

func (userServ *userService) FindByPassEmail(email, password string) (entity.AuthUser, error) {
	return userServ.userRepo.FindByPassEmail(email, password)
}
