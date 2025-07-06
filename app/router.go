package app

import (
	"hermawansafrin/belajar-golang-restful-api/controller"
	"hermawansafrin/belajar-golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryConroller controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryConroller.FindAll)
	router.POST("/api/categories", categoryConroller.Create)
	router.PUT("/api/categories/:categoryId", categoryConroller.Update)
	router.DELETE("/api/categories/:categoryId", categoryConroller.Delete)
	router.GET("/api/categories/:categoryId", categoryConroller.FindById)

	router.PanicHandler = exception.ErrorHandler

	return router
}
