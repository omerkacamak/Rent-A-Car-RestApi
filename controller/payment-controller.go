package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/service"
)

type PaymentController interface {
	GetPaymentByOrderId(ctx *gin.Context)
	GetPaymentsWithOrder(ctx *gin.Context)
}

type paymentController struct {
	service service.PaymentService
}

func NewPaymentController() PaymentController {
	return &paymentController{
		service: service.NewPaymentService(),
	}
}

func (paymentCtrl *paymentController) GetPaymentByOrderId(ctx *gin.Context) {
	orderId := ctx.Param("id")

	orderIdint, err := strconv.Atoi(orderId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err2 := paymentCtrl.service.GetPaymentByOrderId(orderIdint)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}
	ctx.JSON(http.StatusOK, payment)
}

func (paymentCtrl *paymentController) GetPaymentsWithOrder(ctx *gin.Context) {
	payments, err := paymentCtrl.service.GetPaymentsWithOrder()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payments)

}
