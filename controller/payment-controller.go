package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/service"
)

type PaymentController interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)

	GetPaymentByOrderId(ctx *gin.Context)
}

type paymentController struct {
	service service.PaymentService
}

func NewPaymentController() PaymentController {
	return &paymentController{
		service: service.NewPaymentService(),
	}
}

func (paymentCtrl *paymentController) Save(ctx *gin.Context) {
	var payment entity.Payment

	err := ctx.ShouldBindJSON(&payment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = paymentCtrl.service.Save(payment)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payment)

}

func (paymentCtrl *paymentController) Update(ctx *gin.Context) {
	var payment entity.Payment
	err := ctx.ShouldBindJSON(&payment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = paymentCtrl.service.Update(payment)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payment)
}
func (paymentCtrl *paymentController) Delete(ctx *gin.Context) {
	var payment entity.Payment
	err := ctx.ShouldBindJSON(&payment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = paymentCtrl.service.Delete(payment)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

func (paymentCtrl *paymentController) FindAll(ctx *gin.Context) {
	payments := paymentCtrl.service.FindAll()

	ctx.JSON(http.StatusOK, payments)

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, payment)
}
