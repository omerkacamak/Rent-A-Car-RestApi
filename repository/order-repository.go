package repository

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Save(order entity.Order) error
	Update(order entity.Order) error
	Delete(order entity.Order) error
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
func (orderRepo *orderRepository) Save(order entity.Order) error {
	result := orderRepo.connection.Create(&order)

	if result.Error != nil {
		return result.Error
	}
	return nil

}
func (orderRepo *orderRepository) Update(order entity.Order) error {
	result := orderRepo.connection.Save(&order)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (orderRepo *orderRepository) Delete(order entity.Order) error {
	result := orderRepo.connection.Delete(&order)

	if result.Error != nil {
		return result.Error
	}
	return nil
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
