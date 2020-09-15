package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping checks if server is up
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
