package app

import "github.com/pt-abhishek/users-api/controllers"

func mapURLS() {
	router.GET("/ping", controllers.Ping)

	router.POST("/users", controllers.CreateUser)
	router.GET("/users/search", controllers.SearchUser)
	router.GET("/user/:user_id", controllers.FindUser)
	router.PUT("/user/:user_id", controllers.UpdateUser)
	router.PATCH("/user/:user_id", controllers.UpdateUser)
	router.DELETE("/user/:user_id", controllers.DeleteUser)

}
