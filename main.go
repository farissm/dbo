package main

import (
	"fmt"
	"os"
	"net/http"

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

	router.POST("/dbo/login", dboController.Login)

	router.POST("/dbo/create-customer", dboController.CreateCustomer)
	// router.POST("/api/sendPefindoInquiry", auth.AuthMiddleware(authService, userService), pefindoHandler.SendPefindoUnderlying)
	// router.GET("/api/readPefindoInquiry", auth.AuthMiddleware(authService, userService), pefindoHandler.ReadPefindoUnderlying)

	router.Run(":"+os.Getenv("APP_PORT"))
}