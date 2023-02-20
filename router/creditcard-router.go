package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
)

func CreditCardRouter(router *gin.RouterGroup) {
	ctrl := controller.NewCreditCardController()
	router.GET("/", ctrl.FindAll)
	router.POST("/", ctrl.Save)
	router.PATCH("/", ctrl.Update)
	router.DELETE("/", ctrl.Delete)
}