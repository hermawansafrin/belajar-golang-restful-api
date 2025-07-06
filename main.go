package main

import (
	"hermawansafrin/belajar-golang-restful-api/app"
	"hermawansafrin/belajar-golang-restful-api/controller"
	"hermawansafrin/belajar-golang-restful-api/helper"
	"hermawansafrin/belajar-golang-restful-api/middleware"
	"hermawansafrin/belajar-golang-restful-api/repository"
	"hermawansafrin/belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryConroller := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryConroller)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
