package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
)

func OrderRouter(router *gin.RouterGroup) {
	ctrl := controller.NewOrderController()

	router.GET("/", ctrl.FindAll)
	router.POST("/", ctrl.Save)
	router.PUT("/", ctrl.Save)
	router.DELETE("/", ctrl.Delete)
	router.GET("/withcustomer", ctrl.GetAllWithCustomer)
}
