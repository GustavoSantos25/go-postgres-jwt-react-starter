package router

import (
	"github.com/rogaha/go-postgres-jwt-react-starter/server/controller"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter setup routing here
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORSMiddleware())

	// routes
	router.GET("/ping", controller.Pong)
	router.POST("/register", controller.Create)
	router.POST("/login", controller.Login)
	router.GET("/session", controller.Session)
	router.POST("/customers", controller.CreateCustomer)
	router.GET("/customers", controller.GetAllCustomers)
	router.GET("/customers/:id", controller.GetCustomerById)
	router.POST("/paymentmethods", controller.CreatePaymentMethod)
	router.GET("/paymentmethods", controller.GetAllPaymentMethods)
	router.GET("/paymentmethods/:id", controller.GetPaymentMethodById)
	return router
}
