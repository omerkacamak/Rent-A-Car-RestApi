package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
	"github.com/omerkacamak/rentacar-golang/entity"
	"github.com/omerkacamak/rentacar-golang/middleware"
	"github.com/omerkacamak/rentacar-golang/service"
)

var (
	diziVehicle       []entity.Vehicle
	vehicleService    service.VehicleService       = service.NewVehicleService()
	VehicleController controller.VehicleController = controller.NewVehicleController(vehicleService)

	CustomerController controller.CustomerController = controller.NewCustomerController()
	OrderController    controller.OrderController    = controller.NewOrderController()
)

func main() {

	server := gin.New()

	middleware.AutoMigrate()

	server.Use(gin.Recovery())

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "merhaba",
		})
	})

	apiRoutes := server.Group("/vehicle")
	{
		apiRoutes.POST("/", func(ctx *gin.Context) {
			err := VehicleController.Save(ctx)
			if err != nil {
				println("burda hataa")
				ctx.JSON(http.StatusBadRequest, gin.H{

					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "başarılı"})
			}

		})

		apiRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, VehicleController.FindAll())
		})
	}

	apiRoutes2 := server.Group("/customer")
	{
		apiRoutes2.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, CustomerController.FindAll())
		})

		apiRoutes2.POST("/", func(ctx *gin.Context) {
			customer, err := CustomerController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "başarılı",
					"customer": customer,
				})
			}
		})
	}

	apiRoutes3 := server.Group("/order")
	{
		//get
		apiRoutes3.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, OrderController.FindAll())
		})

		//post
		apiRoutes3.POST("/", func(ctx *gin.Context) {
			order, err := OrderController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "başarılı",
					"order":   order,
				})
			}

		})

		// withcustomer
		apiRoutes3.GET("/withCustomer", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, OrderController.GetAllWithCustomer())
		})
	}
	//middleware.AutoMigrate()

	server.Run(":8080")
}
