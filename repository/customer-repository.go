package repository

import (
	"gorm.io/gorm"
)

type CustomerRepository interface {
	// Save(customer entity.Customer) error
	// Update(customer entity.Customer) error
	// Delete(customer entity.Customer) error
	// FindAll() []entity.Customer

	// GENERIC DIŞI METHOT VARSA BURAYA
}

type customerRepository struct {
	connection *gorm.DB
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{
		connection: DB,
	}
}
