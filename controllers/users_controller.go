package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
		return
	}
	c.JSON(http.StatusCreated, result)
}

//SearchUser elastic search i guess
func SearchUser(c *gin.Context) {

}

//FindUser finds a user by ID
func FindUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		parseErr := utils.NewBadRequestError("Invalid user id in params")
		c.JSON(parseErr.Code, parseErr)
		return
	}
	result, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
