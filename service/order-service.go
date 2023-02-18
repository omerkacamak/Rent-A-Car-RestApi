package service

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type OrderService interface {
	Save(customer entity.Order) error
	Update(customer entity.Order) error
	Delete(customer entity.Order) error
	FindAll() []entity.Order //return tipi

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

func (ord *orderService) Save(order entity.Order) error {

	err := ord.orderRepository.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (ord *orderService) Update(order entity.Order) error {
	err := ord.orderRepository.Update(order)
	if err != nil {
		return err
	}
	return nil
}

func (ord *orderService) Delete(order entity.Order) error {
	err := ord.orderRepository.Delete(order)
	if err != nil {
		return err
	}
	return nil
}

func (ord *orderService) FindAll() []entity.Order {
	return ord.orderRepository.FindAll()
}

func (ord *orderService) GetAllWithCustomer() []entity.Order {
	return ord.orderRepository.GetAllWithCustomer()
}
