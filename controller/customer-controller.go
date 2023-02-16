package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/service"
)

type CustomerController interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type customerController struct {
	service service.CustomerService
}

func NewCustomerController() CustomerController {

	return &customerController{
		service: service.NewCustomerService(),
	}
}

func (cst *customerController) Save(ctx *gin.Context) {
	var customer entity.Customer
	err := ctx.ShouldBindJSON(&customer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cst.service.Save(customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)

}
func (cst *customerController) Update(ctx *gin.Context) {
	var customer entity.Customer
	err := ctx.ShouldBindJSON(&customer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cst.service.Update(customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)

}

func (cst *customerController) Delete(ctx *gin.Context) {
	var customer entity.Customer
	err := ctx.ShouldBindJSON(&customer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cst.service.Delete(customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)

}

func (cst *customerController) FindAll(ctx *gin.Context) {
	customers := cst.service.FindAll()

	ctx.JSON(http.StatusOK, customers)
}
