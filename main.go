package main

import (
	"tosshy-blog-api/controller"
	"tosshy-blog-api/db"
	"tosshy-blog-api/repository"
	"tosshy-blog-api/router"
	"tosshy-blog-api/usecase"
	"tosshy-blog-api/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
