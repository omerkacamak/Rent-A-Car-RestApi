package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
	"github.com/omerkacamak/rentacar-golang/entity"
)

func UserRouter(router *gin.RouterGroup) {
	genCtrl := controller.NewGenericController[entity.AuthUser]()
	router.GET("/", genCtrl.FindAll)
	router.POST("/", genCtrl.Save)
	router.PATCH("/:id", genCtrl.Update)
	router.DELETE("/:id", genCtrl.Delete)
	router.GET("/:id", genCtrl.GetById)
}
