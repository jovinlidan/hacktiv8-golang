package httpserver

import (
	"clean-arch-1/httpserver/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	// fmt.Print("test")
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)

	return router
}
