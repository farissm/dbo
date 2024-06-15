package controller

import (
	"gorm.io/gorm"
	"net/http"

	"dbo/model"
	"dbo/helper"

	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (c *Controller) CreateCustomer(ctx *gin.Context) {
	var customer model.Customer
	var customerCheck model.Customer

	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		log.Errorf("FAILED bind json input data. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusBadRequest, "Failed to create account. Please call customer service")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.db.Where("email = ?", customer.Email).First(&customerCheck).Error
	if err != nil {
		log.Errorf("FAILED search existing email. Error: %s", err.Error())
	}

	if customer.Email == customerCheck.Email {
		response := helper.APIResponse(http.StatusBadRequest, "your email is already exist")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	hashedPassword, err := helper.HashPassword(customer.Password)
	if err != nil {
		log.Errorf("FAILED hash password. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusBadRequest, "Failed to create account. Please call customer service")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	customer.Password = string(hashedPassword)

	result := c.db.Create(&customer)
	if result.Error != nil {
		log.Errorf("FAILED when insert data to db. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusBadRequest, "Failed to create account. Please call customer service")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := helper.GenerateToken(customer.Username)
	if err != nil {
		log.Errorf("FAILED GenerateToken. Error= %s", err.Error())

		response := helper.APIResponse(http.StatusBadRequest, "Failed create account. Please call customer service")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.LoginAPIResponse(token.SignedToken, int(token.ExpiresIn), http.StatusCreated)
	ctx.JSON(http.StatusCreated, response)
	return
}

