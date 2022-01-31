package routes

import (
	"github.com/coffeedragon96/go-jwt-authentication/controllers"
	controller "github.com/coffeedragon96/go-jwt-authentication/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
