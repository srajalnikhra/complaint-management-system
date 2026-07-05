package routes

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/controllers"
	"github.com/srajalnikhra/complaint-management-system/internal/middleware"
)

func RegisterComplaintRoutes() {

	http.Handle("/complaints", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateComplaint)))

	http.Handle("/my-complaints", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetMyComplaints)))

	http.Handle("/complaints/", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			controllers.GetComplaintByID(w, r)

		case http.MethodPut:
			controllers.UpdateComplaint(w, r)

		case http.MethodDelete:
			controllers.DeleteComplaint(w, r)

		default:
			http.NotFound(w, r)
		}
	})))
}
