package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/service"
)

type OrderController interface {
	Save(ctx *gin.Context) (entity.Order, error)
	// Update(vehicle entity.Vehicle) entity.Vehicle
	// Delete(vehicle entity.Vehicle) bool
	FindAll() []entity.Order

	GetAllWithCustomer() []entity.Order
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

func (ordrct *orderController) Save(ctx *gin.Context) (entity.Order, error) {
	var order entity.Order
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		return order, err
	}
	ordrct.service.Save(order)
	return order, nil
}

func (ordrct *orderController) FindAll() []entity.Order {
	return ordrct.service.FindAll()
}

func (ordrct *orderController) GetAllWithCustomer() []entity.Order {
	return ordrct.service.GetAllWithCustomer()
}
