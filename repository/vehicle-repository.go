package repository

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	Save(vehicle entity.Vehicle) error
	Update(vehicle entity.Vehicle) error
	Delete(vehicle entity.Vehicle) error
	FindAll() []entity.Vehicle
}

type vehicleRepository struct {
	connection *gorm.DB
}

func NewVehicleRepository() VehicleRepository {
	db, err := NewDbContext()
	if err != nil {
		err.Error()
		println("vehicle repo patladı")
	}
	db.AutoMigrate(&entity.Vehicle{})
	return &vehicleRepository{
		connection: db,
	}
}

func (vehRepo *vehicleRepository) Save(vehicle entity.Vehicle) error {
	result := vehRepo.connection.Create(&vehicle)
	if result.Error != nil {
		println("BÜSBÜYÜK HATA -*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-")
		return result.Error
	}

	return nil

}
func (vehRepo *vehicleRepository) Update(vehicle entity.Vehicle) error {
	result := vehRepo.connection.Save(&vehicle)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (vehRepo *vehicleRepository) Delete(vehicle entity.Vehicle) error {
	result := vehRepo.connection.Delete(&vehicle)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (vehRepo *vehicleRepository) FindAll() []entity.Vehicle {
	var vehicles []entity.Vehicle
	vehRepo.connection.Set("gorm:auto_preload", true).Find(&vehicles)
	return vehicles
}
