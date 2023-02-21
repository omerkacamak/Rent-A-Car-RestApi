package service

import (
	"github.com/omerkacamak/rentacar-golang/repository"
)

type VehicleService interface {
}

type vehicleService struct {
	//vehicle []entity.Vehicle
	vehicleRepo repository.VehicleRepository
}

func NewVehicleService() VehicleService {
	repo := repository.NewVehicleRepository()
	return &vehicleService{
		vehicleRepo: repo,
	}
}
