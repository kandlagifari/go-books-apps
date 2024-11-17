package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kandlagifari/go-books-apps/controllers"
	"github.com/kandlagifari/go-books-apps/middleware"
)

func RegisterBookRoutes(router *gin.Engine) {
	bookGroup := router.Group("/api/books", middleware.AuthMiddleware)
	{
		bookGroup.GET("", controllers.GetBooks)
		bookGroup.POST("", controllers.CreateBook)
		bookGroup.GET("/:id", controllers.GetBookByID)
		bookGroup.DELETE("/:id", controllers.DeleteBook)
		bookGroup.PUT("/:id", controllers.UpdateBook)
	}
}
