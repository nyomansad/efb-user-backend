package routes

import (
	"EFB-User-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User management
	r.POST("/register", controllers.Register)
	r.GET("/users", controllers.GetUsers)
	r.DELETE("/user/:id", controllers.DeleteUser)

	// Airline management
	r.POST("/airlines", controllers.CreateAirline)
	r.GET("/airlines", controllers.GetAirlines)
	r.DELETE("/airline/:id", controllers.DeleteAirline)

        r.POST("/login", controllers.UserLogin)
        r.POST("/logout", controllers.UserLogout)

	return r
}
