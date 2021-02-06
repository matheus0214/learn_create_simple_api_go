package routes

import (
	usersroutes "matheus0214/learn_gin/modules/users/routes"

	"github.com/gin-gonic/gin"
)

// Routes all application routes
func Routes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", usersroutes.NewUserRoutes().Show)
		userGroup.POST("/", usersroutes.NewUserRoutes().Create)
		userGroup.OPTIONS("/", usersroutes.NewUserRoutes().Options)
	}
}
