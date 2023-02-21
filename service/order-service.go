package service

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type OrderService interface {
	GetAllWithCustomer() []entity.Order
}

type orderService struct {
	//orders []entity.Order
	orderRepository repository.OrderRepository
}

func NewOrderService() OrderService {
	repo := repository.NewOrderRepository()
	return &orderService{
		orderRepository: repo,
	}
}

func (ord *orderService) GetAllWithCustomer() []entity.Order {
	return ord.orderRepository.GetAllWithCustomer()
}
