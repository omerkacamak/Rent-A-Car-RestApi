package service

import (
	"github.com/omerkacamak/rentacar-golang/repository"
)

type CustomerService interface {
}

type customerService struct { //struct ile aslında nesnesi
	//customers []entity.Customer
	// repository katmanı gekldiğindfe repository koyulacak
	customers repository.CustomerRepository
}

func NewCustomerService() CustomerService {

	return &customerService{
		customers: repository.NewCustomerRepository(),
	}
}
