package controller

import (
	"hermawansafrin/belajar-golang-restful-api/helper"
	"hermawansafrin/belajar-golang-restful-api/model/web"
	"hermawansafrin/belajar-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService // ngk perlu pointer kalo ini interface
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// read json first
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId, err := strconv.Atoi(params.ByName("categoryId")) // convert string ke integer
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = categoryId // convert ke integer dulu

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId")) // convert string ke integer
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId")) // convert string ke integer
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
