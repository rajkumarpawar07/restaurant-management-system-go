package routes

import (
	"restaurent-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	// incomingRoutes.PUT("/users/:id", controllers.UpdateUser())
	// incomingRoutes.DELETE("/users/:id", controllers.DeleteUser())
	// incomingRoutes.POST("users/logout", controllers.Logout())
	// incomingRoutes.POST("users/refresh", controllers.refresh())


}