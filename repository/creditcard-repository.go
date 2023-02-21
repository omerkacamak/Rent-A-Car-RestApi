package repository

import (
	"gorm.io/gorm"
)

type CreditCardRepository interface {
	// Save(creditcard entity.CreditCard) error
	// Update(creditcard entity.CreditCard) error
	// Delete(creditcard entity.CreditCard) error
	// FindAll() []entity.CreditCard

	//GetById(id int) (*entity.CreditCard, error)
}

type creditCardRepository struct {
	connection *gorm.DB
}

func NewCreditCardRepository() CreditCardRepository {
	return &creditCardRepository{
		connection: DB,
	}
}
