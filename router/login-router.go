package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
)

func LoginRouter(router *gin.RouterGroup) {
	ctrl := controller.NewLoginController()
	router.GET("/:email/:password", ctrl.GetToken)

}
