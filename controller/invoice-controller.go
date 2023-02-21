package controller

import (
	"github.com/omerkacamak/rentacar-golang/service"
)

type InvoiceController interface {
}

type invoiceController struct {
	service service.InvoiceService
}

func NewInvoiceController() InvoiceController {
	return &invoiceController{
		service: service.NewInvoiceService(),
	}
}
