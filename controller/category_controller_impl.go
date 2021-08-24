package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		categoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryCreateRequest := web.CategoryCreateRequest{}

	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.categoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// writter.Header().Add("Content-Type", "Application/json")
	// encoder := json.NewEncoder(writter)
	// err := encoder.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(writter, webResponse)
}
func (controller *CategoryControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}

	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	categoryUpdateRequest.Id = id
	helper.PanicIfError(err)

	categoryResponse := controller.categoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// writter.Header().Add("Content-Type", "Application/json")
	// encoder := json.NewEncoder(writter)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(writter, webResponse)
}
func (controller *CategoryControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.categoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	// writter.Header().Add("Content-Type", "Application/json")
	// encoder := json.NewEncoder(writter)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(writter, webResponse)
}
func (controller *CategoryControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.categoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	writter.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *CategoryControllerImpl) FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryResponses := controller.categoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	// writter.Header().Add("Content-Type", "Application/json")
	// encoder := json.NewEncoder(writter)
	// err := encoder.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(writter, webResponse)
}
