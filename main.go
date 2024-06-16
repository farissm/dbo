package main

import (
	"fmt"
	"os"
	"net/http"

	"dbo/auth"
	"dbo/config"
	"dbo/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Starting DBO test ...")
	
	rl, err := rotatelogs.New("./storage/logs/dbo-%Y-%m-%d.log")
	if err != nil {
		panic(err)
	}

	log.SetOutput(rl)
	log.Info("Success logging into file")

	db := config.ConnectDB()
	fmt.Println("Connected to database")

	log.Info(db)

	dboController := controller.NewController(db)

	router := gin.Default()

	router.GET("/", func(ct *gin.Context) {
		ct.JSON(http.StatusOK, "Test Backend Enginer DBO")
	})

	dbo := router.Group("/dbo")

	dbo.POST("/login", dboController.Login)


	customers := dbo.Group("/customers")
	customers.POST("/create-customer", dboController.CreateCustomer)
	customers.GET("/get-all-customer", auth.AuthMiddleware(db), dboController.GetCustomers)
	customers.GET("/get-customer-by-id", auth.AuthMiddleware(db), dboController.GetCustomerById)
	customers.PUT("/update-customer", auth.AuthMiddleware(db), dboController.UpdateCustomer)
	customers.DELETE("/delete-customer", auth.AuthMiddleware(db), dboController.DeleteCustomer)

	orders := dbo.Group("/orders")
	orders.POST("/create-order", auth.AuthMiddleware(db), dboController.CreateOrder)
	orders.GET("/get-all-order", auth.AuthMiddleware(db), dboController.GetOrders)
	orders.GET("/get-order-by-id", auth.AuthMiddleware(db), dboController.GetOrderById)
	orders.PUT("/update-order", auth.AuthMiddleware(db), dboController.UpdateOrder)
	orders.DELETE("/delete-order", auth.AuthMiddleware(db), dboController.DeleteOrder)

	router.Run(":"+os.Getenv("APP_PORT"))
}