package repository

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"gorm.io/gorm"
)

type CreditCardRepository interface {
	Save(creditcard entity.CreditCard) error
	Update(creditcard entity.CreditCard) error
	Delete(creditcard entity.CreditCard) error
	FindAll() []entity.CreditCard
}

type creditCardRepository struct {
	connection *gorm.DB
}

func NewCreditCardRepository() CreditCardRepository {
	return &creditCardRepository{
		connection: DB,
	}
}

func (creditRepo *creditCardRepository) Save(creditcard entity.CreditCard) error {
	err := creditRepo.connection.Create(&creditcard)

	if err.Error != nil {
		return err.Error
	}
	return nil
}
func (creditRepo *creditCardRepository) Update(creditcard entity.CreditCard) error {
	err := creditRepo.connection.Save(&creditcard)

	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (creditRepo *creditCardRepository) Delete(creditcard entity.CreditCard) error {
	err := creditRepo.connection.Delete(&creditcard)

	if err.Error != nil {
		return err.Error
	}
	return nil
}
func (creditRepo *creditCardRepository) FindAll() []entity.CreditCard {
	var creditcards []entity.CreditCard
	creditRepo.connection.Set("gorm:auto_preload", true).Find(&creditcards)

	return creditcards

}
