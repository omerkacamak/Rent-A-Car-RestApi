package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/service"
)

type LoginController interface {
	GetToken(ctx *gin.Context)
}

type loginController struct {
	service service.AuthService
}

func NewLoginController() LoginController {
	println("login controller olustu ********************************")
	return &loginController{
		service: service.NewAuthService(),
	}
}

func (loginCtr *loginController) GetToken(ctx *gin.Context) {
	email := ctx.Param("email")
	password := ctx.Param("password")

	token, err := loginCtr.service.GenerateToken(email, password)
	if err != nil {
		println("get token controller hatası")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}

}
