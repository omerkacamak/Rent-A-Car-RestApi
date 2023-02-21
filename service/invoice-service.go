package service

import (
	"github.com/omerkacamak/rentacar-golang/repository"
)

type InvoiceService interface {
}

type invoiceService struct {
	repo repository.InvoiceRepository
}

func NewInvoiceService() InvoiceService {
	return &invoiceService{
		repo: repository.NewInvoiceRepository(),
	}
}
