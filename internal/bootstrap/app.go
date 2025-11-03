package bootstrap

import (
	"book-api/config"
	"book-api/internal/controllers"
	"book-api/internal/repository"
	"book-api/internal/routes"
	"book-api/internal/services"
	"log"
)

func InitApp() {
	db := config.InitDB()

	categoryRepo := repository.NewCategoryRepo(db)
	bookRepo := repository.NewBookRepo(db)
	userRepo := repository.NewUserRepo(db)

	categorySvc := services.NewCategoryService(categoryRepo)
	bookSvc := services.NewBookService(bookRepo)
	userSvc := services.NewUserService(userRepo)

	categoryCtrl := controllers.NewCategoryController(categorySvc)
	bookCtrl := controllers.NewBookController(bookSvc)
	userCtrl := controllers.NewUserController(userSvc)

	r := routes.SetupRouter(categoryCtrl, bookCtrl, userCtrl)

	log.Println("Server running di port 8080")
	r.Run(":8080")
}
