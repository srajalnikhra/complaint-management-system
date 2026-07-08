package routes

func RegisterRoutes() {
	RegisterUserRoutes()
	RegisterProtectedRoutes()
	RegisterComplaintRoutes()
	RegisterAdminRoutes()
	RegisterHealthRoutes()
}