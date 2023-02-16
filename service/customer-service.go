package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type CustomerService interface {
	Save(customer entity.Customer) error
	Update(customer entity.Customer) error
	Delete(customer entity.Customer) error
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
	//validate.RegisterValidation("lent", myvalidator.CustomerValidator)

	return &customerService{
		customers: repository.NewCustomerRepository(),
	}
}

func (cstmr *customerService) Save(customer entity.Customer) error {

	// err2 := myvalidator.CustomerValidator2(validate, &customer)
	// if err2 != nil {
	// 	println("--> " + "cUSTOMNER VALİ 2 DE HATA VAR")
	// 	return err2
	// }
	cstmr.customers.Save(customer)
	return nil
}

func (cstmr *customerService) Update(customer entity.Customer) error {
	err := cstmr.customers.Update(customer)
	if err != nil {
		return err
	}
	return nil
}
func (cstmr *customerService) Delete(customer entity.Customer) error {
	err := cstmr.customers.Delete(customer)
	if err != nil {
		return err
	}
	return nil
}

func (cstmr *customerService) FindAll() []entity.Customer {
	return cstmr.customers.FindAll()
}
