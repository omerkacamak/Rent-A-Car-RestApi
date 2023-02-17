package service

import (
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/repository"
)

type VehicleService interface {
	Save(vehicle entity.Vehicle) error
	Update(vehicle entity.Vehicle) error
	Delete(vehicle entity.Vehicle) error
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
		println("servise de düştü")
		return err
	}
	return nil
}

func (vehServ *vehicleService) Update(vehicles entity.Vehicle) error {
	//vehServ.vehicle = append(vehServ.vehicle, vehicles)
	err := vehServ.vehicleRepo.Update(vehicles)
	if err != nil {
		return err
	}
	return nil
}

func (vehServ *vehicleService) Delete(vehicles entity.Vehicle) error {
	//vehServ.vehicle = append(vehServ.vehicle, vehicles)
	err := vehServ.vehicleRepo.Delete(vehicles)
	if err != nil {
		return err
	}
	return nil
}

func (vehServ *vehicleService) FindAll() []entity.Vehicle {
	return vehServ.vehicleRepo.FindAll()
}
