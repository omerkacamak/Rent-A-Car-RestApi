package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omerkacamak/rentacar-golang/controller"
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

	//swagger ekleme
	//server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "merhaba",
		})
	})

	//ROUTERS

	myapi := server.Group("/api" /*, middleware.AuthorizeJWT()*/)
	{
		router.CustomerRouter(myapi.Group("/customer"))
		router.UserRouter(myapi.Group("/user"))
		router.VehicleRouter(myapi.Group("/vehicle"))
		router.OrderRouter(myapi.Group("/order"))
		router.CreditCardRouter(myapi.Group("/creditcard"))
		router.PaymentRouter(myapi.Group("/payment"))

	}
	router.LoginRouter(server.Group("/gettoken")) // TOKEN ALINAN ADRESSS

	//ROUTERS

	server.Run(":8080")
}
