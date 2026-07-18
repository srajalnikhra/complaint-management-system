package routes

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/controllers"
)

// RegisterHealthRoutes registers public health diagnostic checks endpoints.
func RegisterHealthRoutes() {
	// Map the health check route.
	http.HandleFunc("/health", controllers.HealthCheck)
}
