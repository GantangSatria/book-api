package routes

import (
	"book-api/internal/controllers"
	"book-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	categoryCtrl *controllers.CategoryController,
	bookCtrl *controllers.BookController,
	userCtrl *controllers.UserController,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/users/register", userCtrl.Register)
		api.POST("/users/login", userCtrl.Login)

		protected := api.Group("", middleware.JWTAuthMiddleware())
		{
			protected.GET("/categories", categoryCtrl.GetAll)
			protected.POST("/categories", categoryCtrl.Create)
			protected.GET("/categories/:id", categoryCtrl.GetByID)
			protected.DELETE("/categories/:id", categoryCtrl.Delete)
			protected.GET("/categories/:id/books", bookCtrl.GetByCategory)

			protected.GET("/books", bookCtrl.GetAll)
			protected.POST("/books", bookCtrl.Create)
			protected.GET("/books/:id", bookCtrl.GetByID)
			protected.DELETE("/books/:id", bookCtrl.Delete)
		}
	}
	return r
}
