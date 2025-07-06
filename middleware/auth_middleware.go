package middleware

import (
	"hermawansafrin/belajar-golang-restful-api/helper"
	"hermawansafrin/belajar-golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-KEY") == "RAHASIA" {
		// ok
		middleware.Handler.ServeHTTP(writer, request) // lanjutkan
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   "X-API-KEY is required",
			// Message: err.(error).Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		// err
	}
}
