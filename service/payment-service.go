package service

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type PaymentService interface {
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
