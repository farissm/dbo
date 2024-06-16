package controller

import (
	"fmt"
	"time"
	"strconv"
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

		response := helper.APIResponse(http.StatusInternalServerError, "Failed to create account. Please call customer service")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	err = c.db.Where("email = ?", customer.Email).Or("username = ?", customer.Username).First(&customerCheck).Error
	if err != nil {
		log.Errorf("FAILED search existing email. Error: %s", err.Error())
	}

	if customer.Email == customerCheck.Email {
		response := helper.APIResponse(http.StatusBadRequest, "your email is already exist")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if customer.Username == customerCheck.Username {
		response := helper.APIResponse(http.StatusBadRequest, "your username is already exist")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}


	hashedPassword, err := helper.HashPassword(customer.Password)
	if err != nil {
		log.Errorf("FAILED hash password. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "Failed to create account. Please call customer service")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	
	customer.Password = string(hashedPassword)

	result := c.db.Create(&customer)
	if result.Error != nil {
		log.Errorf("FAILED when insert data to db. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "Failed to create account. Please call customer service")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := helper.GenerateToken(customer.Username)
	if err != nil {
		log.Errorf("FAILED GenerateToken. Error= %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "Failed create account. Please call customer service")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.LoginAPIResponse(token.SignedToken, int(token.ExpiresIn), http.StatusCreated)
	ctx.JSON(http.StatusCreated, response)
	return
}

func (c *Controller) GetCustomers(ctx *gin.Context) {
	var customers []model.CustomerResponse

	email := ctx.Query("email")
	page := ctx.Query("page")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Errorf("Error when converting page value. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	var limitPage int = 3
	var offsetPage int = (pageInt - 1) * limitPage
	sql := `SELECT * FROM customers WHERE email LIKE ?
	LIMIT ? OFFSET ?`

	if err := c.db.Raw(sql, fmt.Sprintf(`%%%s%%`, email), limitPage, offsetPage).Find(&customers).Error; err != nil {
		log.Errorf("Error get customers. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, customers)
	return
}

func (c *Controller) GetCustomerById(ctx *gin.Context) {
	var customer model.CustomerResponse

	customerId := ctx.Query("custId")

	if customerId == "" {
		response := helper.APIResponse(http.StatusBadRequest, "missing customer id")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	sql := `SELECT * FROM customers
		WHERE id = ?`
	if err := c.db.Raw(sql, customerId).Find(&customer).Error; err != nil {
		log.Errorf("Error get customer. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, customer)
	return
}

func (c *Controller) UpdateCustomer(ctx *gin.Context) {
	var customer model.Customer

	customerId 	:= ctx.Query("custId")
	
	if customerId == "" {
		response := helper.APIResponse(http.StatusBadRequest, "missing customer id")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		log.Errorf("FAILED bind json input data. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "Internal Server Error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	sql := `UPDATE customers SET username = COALESCE(NULLIF(?, ''), username), first_name = COALESCE(NULLIF(?, ''), first_name), last_name = COALESCE(NULLIF(?, ''), last_name), email = COALESCE(NULLIF(?, ''), email), address = COALESCE(NULLIF(?, ''), address), updated_at = ? WHERE id = ?`

	if err = c.db.Exec(sql, customer.Username, customer.FirstName, customer.LastName, customer.Email, customer.Address, time.Now().Format("2006-01-02 15:04:05"), customerId).Error; err != nil {
		log.Errorf("Error update pefindo underlying status at repository %v", err)

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, "Success update data")
	ctx.JSON(http.StatusOK, response)
	return
}

func (c *Controller) DeleteCustomer(ctx *gin.Context) {
	customerId := ctx.Query("custId")

	if customerId == "" {
		response := helper.APIResponse(http.StatusBadRequest, "missing customer id")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := c.db.Unscoped().Delete(&model.Customer{}, customerId).Error; err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, "Internal Server Error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
    }

	response := helper.APIResponse(http.StatusOK, "User deleted")
	ctx.JSON(http.StatusOK, response)
	return
}