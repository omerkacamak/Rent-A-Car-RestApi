package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/service"
)

type VehicleController interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
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

func (vehController *vehicleController) Save(ctx *gin.Context) {
	var vehicle entity.Vehicle
	err := ctx.ShouldBindJSON(&vehicle)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = vehController.service.Save(vehicle)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vehicle)

}

func (vehController *vehicleController) Update(ctx *gin.Context) {
	var vehicle entity.Vehicle
	err := ctx.ShouldBindJSON(&vehicle)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = vehController.service.Update(vehicle)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vehicle)

}

func (vehController *vehicleController) Delete(ctx *gin.Context) {
	var vehicle entity.Vehicle
	err := ctx.ShouldBindJSON(&vehicle)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = vehController.service.Delete(vehicle)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vehicle)

}

func (vehController *vehicleController) FindAll(ctx *gin.Context) {
	customers := vehController.service.FindAll()

	ctx.JSON(http.StatusOK, customers)

}
