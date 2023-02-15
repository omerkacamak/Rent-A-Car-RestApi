package repository

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Save(order entity.Order)
	Update(order entity.Order)
	Delete(order entity.Order)
	FindAll() []entity.Order

	GetAllWithCustomer() []entity.Order
}

type orderRepository struct {
	connection *gorm.DB
}

func NewOrderRepository() OrderRepository {
	db, err := NewDbContext()
	if err != nil {
		err.Error()
		println("order repo patladı")
	}
	db.AutoMigrate(&entity.Order{})
	return &orderRepository{
		connection: db,
	}
}
func (orderRepo *orderRepository) Save(order entity.Order) {
	orderRepo.connection.Create(&order)
}
func (orderRepo *orderRepository) Update(order entity.Order) {
	orderRepo.connection.Save(&order)
}

func (orderRepo *orderRepository) Delete(order entity.Order) {
	orderRepo.connection.Delete(&order)
}

func (orderRepo *orderRepository) FindAll() []entity.Order {
	var orders []entity.Order
	orderRepo.connection.Set("gorm:auto_preload", true).Find(&orders)
	return orders
}

func (orderRepo *orderRepository) GetAllWithCustomer() []entity.Order {
	var orders []entity.Order
	/*err:=*/ orderRepo.connection.Set("gorm:auto_preload", true).Preload("Customer").Find(&orders) //.Error

	return orders
}
