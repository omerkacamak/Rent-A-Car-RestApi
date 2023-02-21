package controller

import (
	"github.com/omerkacamak/rentacar-golang/service"
)

type UserController interface {
}

type userController struct {
	service service.UserService
}

func NewUserController() UserController {
	println("user controller olustu ********************************")
	return &userController{
		service: service.NewUserService(),
	}
}
