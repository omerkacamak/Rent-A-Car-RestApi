package controller

import (
	"github.com/omerkacamak/rentacar-golang/service"
)

type CustomerController interface {
}

type customerController struct {
	service service.CustomerService
}

func NewCustomerController() CustomerController {
	println("customer controller olustu ********************************")
	return &customerController{
		service: service.NewCustomerService(),
	}
}
