package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
	"github.com/omerkacamak/rentacar-golang/middleware"
	"github.com/omerkacamak/rentacar-golang/repository"
	"github.com/omerkacamak/rentacar-golang/router"
)

var (
	OrderController controller.OrderController = controller.NewOrderController()
	UserController  controller.UserController  = controller.NewUserController()
	LoginController controller.LoginController = controller.NewLoginController()
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

	//ROUTERS

	myapi := server.Group("/api", middleware.AuthorizeJWT())
	{
		router.CustomerRouter(myapi.Group("/customer"))
		router.UserRouter(myapi.Group("/user"))
		router.VehicleRouter(myapi.Group("/vehicle"))

	}
	router.LoginRouter(server.Group("/gettoken")) // TOKEN ALINAN ADRESSS

	//ROUTERS

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

	server.Run(":8080")
}
