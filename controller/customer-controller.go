package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/service"
)

type CustomerController interface {
	Save(ctx *gin.Context) (entity.Customer, error)
	// Update(vehicle entity.Vehicle) entity.Vehicle
	// Delete(vehicle entity.Vehicle) bool
	FindAll() []entity.Customer
}

type customerController struct {
	service service.CustomerService
}

func NewCustomerController() CustomerController {

	return &customerController{
		service: service.NewCustomerService(),
	}
}

func (cst *customerController) Save(ctx *gin.Context) (entity.Customer, error) {
	var customer entity.Customer
	err := ctx.ShouldBindJSON(&customer)

	if err != nil {
		return customer, err
	}

	err2 := cst.service.Save(customer)
	if err2 != nil {
		return customer, err2
	}
	return customer, nil
}

func (cst *customerController) FindAll() []entity.Customer {
	return cst.service.FindAll()
}
