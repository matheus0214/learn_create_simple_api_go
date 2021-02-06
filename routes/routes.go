package routes

import (
	"github.com/gin-gonic/gin"
	usersroutes "github.com/matheus0214/projects/learn_gin/modules/users/routes"
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
