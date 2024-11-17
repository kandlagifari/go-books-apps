package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kandlagifari/go-books-apps/controllers"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/users")
	{
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
	}
}
