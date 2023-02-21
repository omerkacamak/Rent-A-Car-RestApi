package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
	"github.com/omerkacamak/rentacar-golang/entity"
)

func PaymentRouter(router *gin.RouterGroup) {
	genCtrl := controller.NewGenericController[entity.Payment]()
	paymentCtrl := controller.NewPaymentController()
	router.GET("/", genCtrl.FindAll)
	router.POST("/", genCtrl.Save)
	router.PATCH("/:id", genCtrl.Update)
	router.DELETE("/:id", genCtrl.Delete)
	router.GET("/:id", genCtrl.GetById)

	router.GET("/orderid/:id", paymentCtrl.GetPaymentByOrderId) // orderid'ye göre payment getir
	router.GET("/withorder", paymentCtrl.GetPaymentsWithOrder)
}
