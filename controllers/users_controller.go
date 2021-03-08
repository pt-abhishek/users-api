package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pt-abhishek/oAuth-library-go/oauth"
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

	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Code, err)
		return
	}

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
	if oauth.GetCallerID(c.Request) != result.ID {
		c.JSON(http.StatusOK, result.Marshall(true))
		return
	}
	c.JSON(http.StatusOK, result.Marshall(oauth.IsPublic(c.Request)))
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

//Login gets a user which matches the user id and password
func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		parseErr := utils.NewBadRequestError("Unable to parse JSON, invalid request")
		c.JSON(parseErr.Code, parseErr)
		return
	}
	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	//on login create a Oauth token
	var tokenReq = oauth.TokenRequest{
		GrantType:    "client_credentials",
		Scope:        "USERS_API,BOOKS_API,OPENIDCONNECT",
		UserID:       user.ID,
		ClientID:     "2247be5f-56c6-4ec0-bebc-b99b720ede92",
		ClientSecret: "9fc37600-4ec0-4695-9c1f-7f91f4c892fc",
	}
	accessToken, tokenErr := oauth.CreateToken(tokenReq)
	if tokenErr != nil {
		c.JSON(tokenErr.Code, tokenErr)
		return
	}
	c.SetCookie("sessionToken", accessToken.JWT, int(accessToken.Expires), "/", "localhost:8079", false, true)
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
