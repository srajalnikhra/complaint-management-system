package routes

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/controllers"
)

func RegisterHealthRoutes() {
	http.HandleFunc("/health", controllers.HealthCheck)
}