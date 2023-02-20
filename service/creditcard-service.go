package service

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type CreditCardService interface {
	Save(creditcard entity.CreditCard) error
	Update(creditcard entity.CreditCard) error
	Delete(creditcard entity.CreditCard) error
	FindAll() []entity.CreditCard
}

type creditCardService struct {
	repo repository.CreditCardRepository
}

func NewCreditCardService() CreditCardService {
	return &creditCardService{
		repo: repository.NewCreditCardRepository(),
	}
}

func (creditService *creditCardService) Save(creditcard entity.CreditCard) error {
	err := creditService.repo.Save(creditcard)
	if err != nil {
		return err
	}
	return nil
}

func (creditService *creditCardService) Update(creditcard entity.CreditCard) error {
	err := creditService.repo.Update(creditcard)
	if err != nil {
		return err
	}
	return nil
}
func (creditService *creditCardService) Delete(creditcard entity.CreditCard) error {
	err := creditService.repo.Delete(creditcard)
	if err != nil {
		return err
	}
	return nil
}

func (creditService *creditCardService) FindAll() []entity.CreditCard {

	return creditService.repo.FindAll()
}
