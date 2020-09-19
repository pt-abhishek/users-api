package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pt-abhishek/users-api/logger"
)

var (
	router = gin.Default()
)

//StartApplication starts the app
func StartApplication() {
	mapURLS()
	logger.Log.Info("application started on 8081")
	router.Run(":8081")
}
