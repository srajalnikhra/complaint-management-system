package routes

import (
	"net/http"
	"strings"

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

					case http.MethodPatch:
						controllers.UpdateComplaintStatus(w, r)

					default:
						http.NotFound(w, r)
					}
				}),
			),
		),
	)

	http.Handle(
		"/admin/users",
		middleware.AuthMiddleware(
			middleware.AdminMiddleware(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

					switch r.Method {

					case http.MethodGet:
						controllers.GetAllUsers(w, r)

					default:
						http.NotFound(w, r)
					}
				}),
			),
		),
	)

	http.Handle(
		"/admin/users/",
		middleware.AuthMiddleware(
			middleware.AdminMiddleware(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					switch r.Method {

					case http.MethodPatch:

						if strings.HasSuffix(r.URL.Path, "/role") {
							controllers.UpdateUserRole(w, r)
							return
						}

						if strings.HasSuffix(r.URL.Path, "/status") {
							controllers.UpdateUserStatus(w, r)
							return
						}

						http.NotFound(w, r)

					case http.MethodDelete:

						controllers.DeleteUser(w, r)

					default:

						http.NotFound(w, r)
					}
				}),
			),
		),
	)
}
