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

	GetPaymentByOrderId(id int) (*entity.Payment, error) // order id ye göre payment getirme
	GetPaymentsWithOrder() (*[]entity.Payment, error)
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

func (paymentRepo *paymentRepository) GetPaymentByOrderId(id int) (*entity.Payment, error) {
	var payment entity.Payment

	result := paymentRepo.connection.Where("order_id= ?", id).First(&payment)
	//result := paymentRepo.connection.Where(&entity.Payment{OrderID: id}).First(&payment)
	var payo = &entity.Payment{ID: 1, CreditCardID: 1, OrderID: &id}

	if result.Error != nil {
		println("hata var hata hata :: + " + result.Error.Error())

		return payo, result.Error
	}
	return &payment, nil
}

func (paymentRepo *paymentRepository) GetPaymentsWithOrder() (*[]entity.Payment, error) {
	var payments []entity.Payment

	err := paymentRepo.connection.Set("gorm:preload", true).Preload("Order").Find(&payments).Error

	if err != nil {
		return nil, err
	}
	return &payments, nil
}
