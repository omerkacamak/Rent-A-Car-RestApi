package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/service"
)

type OrderController interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)

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

func (ordrct *orderController) Save(ctx *gin.Context) {
	var order entity.Order
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err = ordrct.service.Save(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, order)
}

func (ordrct *orderController) Update(ctx *gin.Context) {
	var order entity.Order
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err = ordrct.service.Update(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, order)
}

func (ordrct *orderController) Delete(ctx *gin.Context) {
	var order entity.Order
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err = ordrct.service.Delete(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, order)
}

func (ordrct *orderController) FindAll(ctx *gin.Context) {
	orders := ordrct.service.FindAll()

	ctx.JSON(http.StatusOK, orders)
}

func (ordrct *orderController) GetAllWithCustomer(ctx *gin.Context) {
	orders := ordrct.service.GetAllWithCustomer()
	ctx.JSON(http.StatusOK, orders)
}
