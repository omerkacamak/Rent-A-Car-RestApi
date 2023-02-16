package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/service"
)

type UserController interface {
	Save(ctx *gin.Context) (entity.AuthUser, error)
	Update(ctx *gin.Context) (entity.AuthUser, error)
	Delete(ctx *gin.Context) (entity.AuthUser, error)
	FindAll() []entity.AuthUser

	FindByPassEmail(email, password string) (entity.AuthUser, error)
	//GetAllWithCustomer() []entity.Order
}

type userController struct {
	service service.UserService
}

func NewUserController() UserController {
	return &userController{
		service: service.NewUserService(),
	}
}

func (userContr *userController) Save(ctx *gin.Context) (entity.AuthUser, error) {
	var user entity.AuthUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return user, err
	}
	userContr.service.Save(user)
	return user, nil
}

func (userContr *userController) Update(ctx *gin.Context) (entity.AuthUser, error) {
	var user entity.AuthUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return user, err
	}
	userContr.service.Update(user)
	return user, nil
}
func (userContr *userController) Delete(ctx *gin.Context) (entity.AuthUser, error) {
	var user entity.AuthUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return user, err
	}
	userContr.service.Update(user)
	return user, nil
}

func (userContr *userController) FindAll() []entity.AuthUser {
	result := userContr.service.FindAll()
	for _, val := range result {
		println("--->usercontroller  " + val.FirstName)
	}
	return result
}

func (userContr *userController) FindByPassEmail(email, password string) (entity.AuthUser, error) {
	return userContr.service.FindByPassEmail(email, password)
}
