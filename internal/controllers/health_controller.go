package controllers

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/config"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

// HealthCheck returns the current health status of the application.
//
// @Summary Health Check
// @Description Check whether the application is running
// @Tags Health
// @Produce json
// @Success 200 {object} utils.APIResponse
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	// Return a success response indicating the backend is running.
	utils.Success(
		w,
		http.StatusOK,
		"Application is healthy",
		map[string]string{
			"status":  "UP",
			"service": config.LoadAppConfig().Name,
			"version": "v1",
		},
	)
}
