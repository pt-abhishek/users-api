package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pt-abhishek/users-api/domain/users"
	"github.com/pt-abhishek/users-api/services"
	"github.com/pt-abhishek/users-api/utils"
)

//CreateUser add new user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("error parsing")
		restErr := utils.NewBadRequestError("Invalid JSON")
		c.JSON(restErr.Code, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		fmt.Println("error while saving")
		c.JSON(saveErr.Code, saveErr)
	}
	c.JSON(http.StatusCreated, result)
}

//SearchUser elastic search i guess
func SearchUser(c *gin.Context) {

}
