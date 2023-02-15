package repository

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(customer entity.Customer)
	Update(customer entity.Customer)
	Delete(customer entity.Customer)
	FindAll() []entity.Customer
}

type customerRepository struct {
	connection *gorm.DB
}

func NewCustomerRepository() CustomerRepository {
	db, err := NewDbContext()
	if err != nil {
		err.Error()
		println("customer repo patladı")
	}
	db.AutoMigrate(&entity.Customer{})
	return &customerRepository{
		connection: db,
	}
}
func (customerRepo *customerRepository) Save(customer entity.Customer) {
	customerRepo.connection.Create(&customer)
}
func (customerRepo *customerRepository) Update(customer entity.Customer) {
	customerRepo.connection.Save(&customer)
}

func (customerRepo *customerRepository) Delete(customer entity.Customer) {
	customerRepo.connection.Delete(&customer)

}

func (customerRepo *customerRepository) FindAll() []entity.Customer {
	var customers []entity.Customer
	customerRepo.connection.Set("gorm:auto_preload", true).Find(&customers)
	return customers
}
