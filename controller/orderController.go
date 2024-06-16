package controller

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	// "io/ioutil"

	"dbo/model"
	"dbo/helper"

	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateOrder(ctx *gin.Context) {
	var order model.Order

	currentUserRaw := ctx.MustGet("currentUser")

    // Assert currentUser to your model.Customer type
    customer, ok := currentUserRaw.(model.Customer)
    if !ok {
		log.Errorf("Error: currentUser is of unexpected type")

		response := helper.APIResponse(http.StatusInternalServerError, "Failed to create order. Please call customer service")
		ctx.JSON(http.StatusInternalServerError, response)
        return
    }

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		log.Errorf("FAILED bind json input data. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "Failed to create order. Please call customer service")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	order.Customer = customer

	result := c.db.Create(&order)
	if result.Error != nil {
		log.Errorf("FAILED when insert data to db. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "Failed to create order. Please call customer service")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusCreated, order)
	return
}

func (c *Controller) GetOrders(ctx *gin.Context) {
	var orders []model.OrderResponse

	email := ctx.Query("name")
	page := ctx.Query("page")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Errorf("Error when converting order page value. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	var limitPage int = 5
	var offsetPage int = (pageInt - 1) * limitPage
	sql := `SELECT * FROM orders WHERE name LIKE ?
	LIMIT ? OFFSET ?`

	if err := c.db.Raw(sql, fmt.Sprintf(`%%%s%%`, email), limitPage, offsetPage).Find(&orders).Error; err != nil {
		log.Errorf("Error get orders. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "Failed fetch data")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, orders)
	return
}

func (c *Controller) GetOrderById(ctx *gin.Context) {
	var order model.OrderResponse

	orderId := ctx.Query("orderId")

	if orderId == "" {
		response := helper.APIResponse(http.StatusBadRequest, "missing order id")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	sql := `SELECT * FROM orders
		WHERE id = ?`
	if err := c.db.Raw(sql, orderId).Find(&order).Error; err != nil {
		log.Errorf("Error get order. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, order)
	return
}

func (c *Controller) UpdateOrder(ctx *gin.Context) {
	var order model.Order

	orderId := ctx.Query("orderId")
	
	if orderId == "" {
		response := helper.APIResponse(http.StatusBadRequest, "missing order id")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		log.Errorf("FAILED bind json input data. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "Internal Server Error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	sql := `UPDATE orders SET name = COALESCE(NULLIF(?, ''), name), quantity = COALESCE(NULLIF(?, 0), quantity), price = COALESCE(NULLIF(?, 0), price), status = COALESCE(NULLIF(?, ''), status), updated_at = ? WHERE id = ?`

	if err = c.db.Exec(sql, order.Name, order.Quantity, time.Now().Format("2006-01-02 15:04:05"), orderId).Error; err != nil {
		log.Errorf("Error update pefindo underlying status at repository %v", err)

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, "Success update order data")
	ctx.JSON(http.StatusOK, response)
	return
}

func (c *Controller) DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Query("orderId")

	if orderId == "" {
		response := helper.APIResponse(http.StatusBadRequest, "missing order id")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := c.db.Unscoped().Delete(&model.Order{}, orderId).Error; err != nil {
		response := helper.APIResponse(http.StatusInternalServerError, "Internal Server Error")
		ctx.JSON(http.StatusInternalServerError, response)
		return
    }

	response := helper.APIResponse(http.StatusOK, "Order deleted")
	ctx.JSON(http.StatusOK, response)
	return
}