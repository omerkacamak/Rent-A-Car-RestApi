package service

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type VehicleService interface {
	Save(vehicle entity.Vehicle) error
	Update(vehicle entity.Vehicle) entity.Vehicle
	Delete(vehicle entity.Vehicle) bool
	FindAll() []entity.Vehicle
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

func (vehServ *vehicleService) Save(vehicles entity.Vehicle) error {
	//vehServ.vehicle = append(vehServ.vehicle, vehicles)
	err := vehServ.vehicleRepo.Save(vehicles)
	if err != nil {
		return err
	}
	return nil
}

func (vehServ *vehicleService) Update(vehicles entity.Vehicle) entity.Vehicle {
	//vehServ.vehicle = append(vehServ.vehicle, vehicles)
	vehServ.vehicleRepo.Update(vehicles)
	return vehicles
}

func (vehServ *vehicleService) Delete(vehicles entity.Vehicle) bool {
	//vehServ.vehicle = append(vehServ.vehicle, vehicles)
	vehServ.vehicleRepo.Delete(vehicles)
	return true
}

func (vehServ *vehicleService) FindAll() []entity.Vehicle {
	return vehServ.vehicleRepo.FindAll()
}
