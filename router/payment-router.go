package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
)

func PaymentRouter(router *gin.RouterGroup) {
	ctrl := controller.NewPaymentController()
	router.GET("/", ctrl.FindAll)
	router.POST("/", ctrl.Save)
	router.PATCH("/", ctrl.Update)
	router.DELETE("/", ctrl.Delete)

	router.GET("/oderid", ctrl.GetPaymentByOrderId)
}
