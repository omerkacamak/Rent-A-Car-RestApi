package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/service"
)

type CreditCardController interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type creditCardController struct {
	service service.CreditCardService
}

func NewCreditCardController() CreditCardController {
	return &creditCardController{
		service: service.NewCreditCardService(),
	}
}

func (creditCtrl *creditCardController) Save(ctx *gin.Context) {
	var creditCard entity.CreditCard
	err := ctx.ShouldBindJSON(&creditCard)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = creditCtrl.service.Save(creditCard)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, creditCard)
}

func (creditCtrl *creditCardController) Update(ctx *gin.Context) {
	var creditCard entity.CreditCard
	err := ctx.ShouldBindJSON(&creditCard)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = creditCtrl.service.Update(creditCard)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, creditCard)
}

func (creditCtrl *creditCardController) Delete(ctx *gin.Context) {
	var creditCard entity.CreditCard
	err := ctx.ShouldBindJSON(&creditCard)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = creditCtrl.service.Delete(creditCard)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, creditCard)
}

func (creditCtrl *creditCardController) FindAll(ctx *gin.Context) {
	creditCards := creditCtrl.service.FindAll()

	ctx.JSON(http.StatusOK, creditCards)
}
