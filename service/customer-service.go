package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type CustomerService interface {
	Save(customer entity.Customer) error
	Update(customer entity.Customer) entity.Customer
	Delete(customer entity.Customer) bool
	FindAll() []entity.Customer //return tipi
}

type customerService struct { //struct ile aslında nesnesi
	//customers []entity.Customer
	// repository katmanı gekldiğindfe repository koyulacak
	customers repository.CustomerRepository
}

var validate *validator.Validate

func NewCustomerService() CustomerService {
	validate = validator.New()
	repo := repository.NewCustomerRepository()
	return &customerService{
		customers: repo,
	}
}

func (cstmr *customerService) Save(customer entity.Customer) error {
	err := validate.Struct(customer)
	if err != nil {
		return err
	}
	for _, err := range err.(validator.ValidationErrors) {
		println("--> " + err.Error())
	}
	cstmr.customers.Save(customer)
	return nil
}

func (cstmr *customerService) Update(customer entity.Customer) entity.Customer {
	cstmr.customers.Update(customer)
	return customer
}
func (cstmr *customerService) Delete(customer entity.Customer) bool {
	cstmr.customers.Delete(customer)
	return true
}

func (cstmr *customerService) FindAll() []entity.Customer {
	return cstmr.customers.FindAll()
}
