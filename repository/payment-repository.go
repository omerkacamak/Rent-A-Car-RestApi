package repository

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Save(payment entity.Payment) error
	Update(payment entity.Payment) error
	Delete(payment entity.Payment) error
	FindAll() []entity.Payment

	GetPaymentByOrderId(id int) (entity.Payment, error)
}

type paymentRepository struct {
	connection *gorm.DB
}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{
		connection: DB,
	}
}

func (paymentRepo *paymentRepository) Save(payment entity.Payment) error {
	err := paymentRepo.connection.Create(&payment)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (paymentRepo *paymentRepository) Update(payment entity.Payment) error {
	err := paymentRepo.connection.Save(&payment)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (paymentRepo *paymentRepository) Delete(payment entity.Payment) error {
	err := paymentRepo.connection.Delete(&payment)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (paymentRepo *paymentRepository) FindAll() []entity.Payment {
	var payments []entity.Payment
	paymentRepo.connection.Set("gorm:auto_preload", true).Find(&payments)
	return payments
}

func (paymentRepo *paymentRepository) GetPaymentByOrderId(id int) (entity.Payment, error) {
	var payment = entity.Payment{OrderID: id}
	result := paymentRepo.connection.First(&payment)
	if result.Error != nil {
		return payment, result.Error
	}
	return payment, nil
}
