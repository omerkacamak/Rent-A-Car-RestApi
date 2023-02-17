package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
)

func UserRouter(router *gin.RouterGroup) {
	ctrl := controller.NewUserController()
	router.GET("/", ctrl.FindAll)
	router.POST("/", ctrl.Save)
	router.PUT("/", ctrl.Save)
	router.DELETE("/", ctrl.Delete)
}
