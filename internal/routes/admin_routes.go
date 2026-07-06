package routes

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/controllers"
	"github.com/srajalnikhra/complaint-management-system/internal/middleware"
)

func RegisterAdminRoutes() {

	http.Handle(
		"/admin/complaints",
		middleware.AuthMiddleware(
			middleware.AdminMiddleware(
				http.HandlerFunc(controllers.GetAllComplaints),
			),
		),
	)

	http.Handle(
		"/admin/complaints/",
		middleware.AuthMiddleware(
			middleware.AdminMiddleware(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

					switch r.Method {

					case http.MethodPut:
						controllers.UpdateComplaintStatus(w, r)

					default:
						http.NotFound(w, r)
					}
				}),
			),
		),
	)
}