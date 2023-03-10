package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
	"github.com/omerkacamak/rentacar-golang/entity"
)

func CustomerRouter(router *gin.RouterGroup) {
	//ctrl := controller.NewCustomerController()
	genCtrl := controller.NewGenericController[entity.Customer]()
	router.GET("/", genCtrl.FindAll)
	router.POST("/", genCtrl.Save)
	router.PATCH("/:id", genCtrl.Update)
	router.DELETE("/:id", genCtrl.Delete)

	router.GET("/:id", genCtrl.GetById)
}
