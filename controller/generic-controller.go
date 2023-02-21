package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/service"
)

type IGenericController interface {
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
}

type GenericController[Tent any] struct {
	genericService service.IGenericService[Tent]
}

func NewGenericController[Tent any]() IGenericController {
	return &GenericController[Tent]{
		genericService: service.NewGenericService[Tent](),
	}
}

func (genericCtrl GenericController[Tent]) Save(ctx *gin.Context) {
	var entity Tent
	err := ctx.ShouldBindJSON(&entity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = genericCtrl.genericService.Save(entity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entity)
}
func (genericCtrl GenericController[Tent]) Update(ctx *gin.Context) {
	var entity Tent

	idStr := ctx.Param("id")
	idint, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err2 := ctx.ShouldBindJSON(&entity)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	err = genericCtrl.genericService.Update(idint, entity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entity)
}
func (genericCtrl GenericController[Tent]) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idint, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := genericCtrl.genericService.GetById(idint)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = genericCtrl.genericService.Delete(*res)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
func (genericCtrl GenericController[Tent]) FindAll(ctx *gin.Context) {
	result, err := genericCtrl.genericService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (genericCtrl GenericController[Tent]) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idint, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := genericCtrl.genericService.GetById(idint)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
