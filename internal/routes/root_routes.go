package routes

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/utils"
)

func registerRootRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		utils.Success(
			w,
			http.StatusOK,
			"Complaint Management System API",
			map[string]string{
				"swagger": "/swagger/index.html",
				"health":  "/health",
			},
		)
	})
}
