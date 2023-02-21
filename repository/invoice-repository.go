package repository

import (
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	// Save(invoice entity.Invoice) error
	// Update(invoice entity.Invoice) error
	// Delete(invoice entity.Invoice) error
	// FindAll() ([]entity.Invoice, error)
}

type invoiceRepository struct {
	connection *gorm.DB
}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{
		connection: DB,
	}
}
