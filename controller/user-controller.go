package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/service"
)

type UserController interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)

	//FindByPassEmail(ctx *gin.Context)
	//GetAllWithCustomer(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController() UserController {
	return &userController{
		service: service.NewUserService(),
	}
}

func (userContr *userController) Save(ctx *gin.Context) {
	var user entity.AuthUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = userContr.service.Save(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (userContr *userController) Update(ctx *gin.Context) {
	var user entity.AuthUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = userContr.service.Update(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
func (userContr *userController) Delete(ctx *gin.Context) {
	var user entity.AuthUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = userContr.service.Delete(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (userContr *userController) FindAll(ctx *gin.Context) {
	users := userContr.service.FindAll()

	ctx.JSON(http.StatusOK, users)

}
