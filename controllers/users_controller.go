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
	userID, err := getID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	result, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
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

	result, updateErr := services.UpdateUser(isPartial, userToUpdate)

	if updateErr != nil {
		c.JSON(updateErr.Code, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

//DeleteUser Deletes a user with the given id
func DeleteUser(c *gin.Context) {
	userID, err := getID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	deletedUser, deleteErr := services.DeleteUser(userID)

	if deleteErr != nil {
		c.JSON(deleteErr.Code, deleteErr)
		return
	}
	c.JSON(http.StatusOK, deletedUser)
}

func getID(id string) (int64, *utils.RestErr) {
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return -1, utils.NewBadRequestError("Invalid user id in params")
	}
	return userID, nil
}
