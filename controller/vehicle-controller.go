package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/service"
)

type VehicleController interface {
	Save(ctx *gin.Context) error
	// Update(vehicle entity.Vehicle) entity.Vehicle
	// Delete(vehicle entity.Vehicle) bool
	FindAll() []entity.Vehicle
}

type vehicleController struct {
	service service.VehicleService
}

func NewVehicleController(service service.VehicleService) VehicleController {
	return &vehicleController{
		service: service,
	}
}

func (vehController *vehicleController) Save(ctx *gin.Context) error {
	var vehicle entity.Vehicle
	err := ctx.ShouldBindJSON(&vehicle)

	if err != nil {
		return err
	}
	err2 := vehController.service.Save(vehicle)
	if err2 != nil {
		return err2
	}
	return nil
}
func (vehController *vehicleController) FindAll() []entity.Vehicle {
	return vehController.service.FindAll()
}
