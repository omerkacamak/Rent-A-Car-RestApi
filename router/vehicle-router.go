package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
)

func VehicleRouter(router *gin.RouterGroup) {

	ctrl := controller.NewVehicleController()

	router.GET("/", ctrl.FindAll)
	router.POST("/", ctrl.Save)
	router.PUT("/", ctrl.Save)
	router.DELETE("/", ctrl.Delete)
}
