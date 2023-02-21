package controller

import (
	"github.com/omerkacamak/rentacar-golang/service"
)

type VehicleController interface {
}

type vehicleController struct {
	service service.VehicleService
}

func NewVehicleController() VehicleController {
	println("vehicle controller olustu ********************************")
	return &vehicleController{
		service: service.NewVehicleService(),
	}
}
