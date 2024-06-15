package controller

import (
	"net/http"

	"dbo/model"
	"dbo/helper"

	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (c *Controller) Login(ctx *gin.Context) {
	username, password, hasAuth := ctx.Request.BasicAuth()
	if ! hasAuth {
		log.Errorf("FAILED hasAuth")

		response := helper.APIResponse(http.StatusUnprocessableEntity, "Login Failed")
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	input := helper.LoginInput{}
	input.Username = username
	input.Password = password

	var customer model.Customer

	err:= c.db.Where("username = ?", username).Find(&customer).Error
	if err != nil {
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Invalid username or password")
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if customer.ID == 0 {
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Invalid username or password")
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if err != nil {
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Invalid username or password")
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := helper.GenerateToken(username)
	if err != nil {
		log.Errorf("FAILED GenerateToken. Error= %s", err.Error())

		response := helper.APIResponse(http.StatusBadRequest, "Login failed")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.LoginAPIResponse(token.SignedToken, int(token.ExpiresIn), http.StatusOK)
	ctx.JSON(http.StatusOK, response)
	return
}