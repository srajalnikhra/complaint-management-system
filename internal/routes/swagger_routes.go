package routes

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterSwaggerRoutes hosts the Swagger API document endpoints.
func RegisterSwaggerRoutes() {
	// Let the swagger module handle doc routing.
	http.Handle(
		"/swagger/",
		httpSwagger.WrapHandler,
	)
}
