package service

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type PaymentService interface {
	Save(payment entity.Payment) error
	Update(payment entity.Payment) error
	Delete(payment entity.Payment) error
	FindAll() []entity.Payment

	GetPaymentByOrderId(id int) (*entity.Payment, error)
	GetPaymentsWithOrder() (*[]entity.Payment, error)
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService() PaymentService {
	return &paymentService{
		repo: repository.NewPaymentRepository(),
	}
}

func (paymentServ *paymentService) Save(payment entity.Payment) error {
	err := paymentServ.repo.Save(payment)
	if err != nil {
		return err
	}
	return nil
}

func (paymentServ *paymentService) Update(payment entity.Payment) error {
	err := paymentServ.repo.Update(payment)
	if err != nil {
		return err
	}
	return nil
}

func (paymentServ *paymentService) Delete(payment entity.Payment) error {
	err := paymentServ.repo.Delete(payment)
	if err != nil {
		return err
	}
	return nil
}

func (paymentServ *paymentService) FindAll() []entity.Payment {
	result := paymentServ.repo.FindAll()

	return result
}

func (paymentServ *paymentService) GetPaymentByOrderId(id int) (*entity.Payment, error) {

	result, err := paymentServ.repo.GetPaymentByOrderId(id)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (paymentServ *paymentService) GetPaymentsWithOrder() (*[]entity.Payment, error) {
	result, err := paymentServ.repo.GetPaymentsWithOrder()
	if err != nil {
		return nil, err
	}
	return result, nil
}
