package service

import (
	"github.com/omerkacamak/rentacar-golang/repository"
)

type CreditCardService interface {
}

type creditCardService struct {
	creditRepo repository.CreditCardRepository
}

func NewCreditCardService() *creditCardService {
	println("hatavaarr")
	return &creditCardService{
		creditRepo: repository.NewCreditCardRepository(),
	}
}
