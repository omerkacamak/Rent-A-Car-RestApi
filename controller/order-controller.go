package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/service"
)

type OrderController interface {
	GetAllWithCustomer(ctx *gin.Context)
}

type orderController struct {
	service service.OrderService
}

func NewOrderController() OrderController {
	println("order controller olustu ********************************")
	return &orderController{
		service: service.NewOrderService(),
	}
}

func (ordrct *orderController) GetAllWithCustomer(ctx *gin.Context) {
	orders := ordrct.service.GetAllWithCustomer()
	ctx.JSON(http.StatusOK, orders)
}
