package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kandlagifari/go-books-apps/controllers"
	"github.com/kandlagifari/go-books-apps/middleware"
)

func RegisterCategoryRoutes(router *gin.Engine) {
	categoryGroup := router.Group("/api/categories", middleware.AuthMiddleware)
	{
		categoryGroup.GET("", controllers.GetCategories)
		categoryGroup.POST("", controllers.CreateCategory)
		categoryGroup.GET("/:id", controllers.GetCategoryByID)
		categoryGroup.DELETE("/:id", controllers.DeleteCategory)
	}
}
