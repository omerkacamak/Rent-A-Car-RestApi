package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/service"
)

type LoginController interface {
	GetToken(ctx *gin.Context) (string, error)
}

type loginController struct {
	service service.AuthService
}

func NewLoginController() LoginController {
	return &loginController{
		service: service.NewAuthService(),
	}
}

func (loginCtr *loginController) GetToken(ctx *gin.Context) (string, error) {
	return "", nil

}
