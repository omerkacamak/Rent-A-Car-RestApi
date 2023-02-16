package repository

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(customer entity.Customer) error
	Update(customer entity.Customer) error
	Delete(customer entity.Customer) error
	FindAll() []entity.Customer
}

type customerRepository struct {
	connection *gorm.DB
}

func NewCustomerRepository() CustomerRepository {
	// db, err := NewDbContext()
	// if err != nil {
	// 	err.Error()
	// 	println("customer repo patladı")
	// }
	// db.AutoMigrate(&entity.Customer{})
	return &customerRepository{
		connection: DB,
	}
}
func (customerRepo *customerRepository) Save(customer entity.Customer) error {
	result := customerRepo.connection.Create(&customer)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
func (customerRepo *customerRepository) Update(customer entity.Customer) error {
	result := customerRepo.connection.Save(&customer)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (customerRepo *customerRepository) Delete(customer entity.Customer) error {
	result := customerRepo.connection.Delete(&customer)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (customerRepo *customerRepository) FindAll() []entity.Customer {
	var customers []entity.Customer
	customerRepo.connection.Set("gorm:auto_preload", true).Find(&customers)
	return customers
}
