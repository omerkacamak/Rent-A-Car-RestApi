package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
	"github.com/omerkacamak/rentacar-golang/middleware"
	"github.com/omerkacamak/rentacar-golang/repository"

	"github.com/omerkacamak/rentacar-golang/service"
)

var (
	vehicleService    service.VehicleService       = service.NewVehicleService()
	VehicleController controller.VehicleController = controller.NewVehicleController(vehicleService)

	CustomerController controller.CustomerController = controller.NewCustomerController()
	OrderController    controller.OrderController    = controller.NewOrderController()
	UserController     controller.UserController     = controller.NewUserController()
	LoginController    controller.LoginController    = controller.NewLoginController()
)

func init() {
	repository.ConnectToDatabase() //burda databaseyi başlattık
}

func main() {

	repository.ConnectToDatabase()

	server := gin.New()

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

	apiRoutes2 := server.Group("/customer", middleware.AuthorizeJWT())
	{
		apiRoutes2.GET("/findall", CustomerController.FindAll)

		apiRoutes2.GET("/", CustomerController.FindAll)

		apiRoutes2.POST("/", CustomerController.Save)
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

	apiRoutes4 := server.Group("/user")
	{
		apiRoutes4.GET("/", UserController.FindAll)

		apiRoutes4.GET("/apikey/:name/:password", LoginController.GetToken)
	}

	apiRoutes5 := server.Group("/login")
	{
		apiRoutes5.GET("/gettoken/:email/:password", LoginController.GetToken)

		server.Run(":8080")
	}
}
