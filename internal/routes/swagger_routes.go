package routes

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterSwaggerRoutes() {
	http.Handle(
		"/swagger/",
		httpSwagger.WrapHandler,
	)
	
}