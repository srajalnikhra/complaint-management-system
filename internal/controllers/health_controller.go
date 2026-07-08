package controllers

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/config"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

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
