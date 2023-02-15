package service

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type OrderService interface {
	Save(customer entity.Order) entity.Order
	Update(customer entity.Order) entity.Order
	Delete(customer entity.Order) bool
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

func (ord *orderService) Save(order entity.Order) entity.Order {
	//ord.orders = append(ord.orders, customer)
	ord.orderRepository.Save(order)
	return order
}

func (ord *orderService) Update(order entity.Order) entity.Order {
	//ord.orders = append(ord.orders, customer)
	ord.orderRepository.Update(order)
	return order
}

func (ord *orderService) Delete(customer entity.Order) bool {
	//ord.orders = append(ord.orders, customer)
	ord.orderRepository.Delete(customer)
	return true
}

func (ord *orderService) FindAll() []entity.Order {
	return ord.orderRepository.FindAll()
}

func (ord *orderService) GetAllWithCustomer() []entity.Order {
	return ord.orderRepository.GetAllWithCustomer()
}
