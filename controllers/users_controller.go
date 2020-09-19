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
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		fmt.Println("error while saving")
		c.JSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//FindUser finds a user by ID
func FindUser(c *gin.Context) {
	userID, err := getID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	result, getErr := services.UserService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//UpdateUser handles update api call
func UpdateUser(c *gin.Context) {
	userID, err := getID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	var userToUpdate users.User
	if err := c.ShouldBindJSON(&userToUpdate); err != nil {
		parseErr := utils.NewBadRequestError("Unable to parse JSON, invalid request")
		c.JSON(parseErr.Code, parseErr)
		return
	}
	userToUpdate.ID = userID
	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UserService.UpdateUser(isPartial, userToUpdate)

	if updateErr != nil {
		c.JSON(updateErr.Code, updateErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//DeleteUser Deletes a user with the given id
func DeleteUser(c *gin.Context) {
	userID, err := getID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	deletedUser, deleteErr := services.UserService.DeleteUser(userID)

	if deleteErr != nil {
		c.JSON(deleteErr.Code, deleteErr)
		return
	}
	c.JSON(http.StatusOK, deletedUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func getID(id string) (int64, *utils.RestErr) {
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return -1, utils.NewBadRequestError("Invalid user id in params")
	}
	return userID, nil
}

//Search gets by searchText
func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserService.Search(status)
	if err != nil {
		c.JSON(err.Code, err)
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
