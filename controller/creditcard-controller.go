package controller

import (
	"github.com/omerkacamak/rentacar-golang/service"
)

type CreditCardController interface {
	// Save(ctx *gin.Context)
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
	// FindAll(ctx *gin.Context)
	// GetById(ctx *gin.Context)
}

type creditCardController struct {
	service service.CreditCardService
}

func NewCreditCardController() CreditCardController {
	return &creditCardController{
		service: service.NewCreditCardService(),
	}
}
