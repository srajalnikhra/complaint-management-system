package routes

// RegisterRoutes registers all route handlers in the application multiplexer.
func RegisterRoutes() {
	// Bind user management and auth routes.
	RegisterUserRoutes()

	// Bind general token-protected profile routes.
	RegisterProtectedRoutes()

	// Bind standard user complaints endpoints.
	RegisterComplaintRoutes()

	// Bind administrative control routes.
	RegisterAdminRoutes()

	// Bind public service health monitoring paths.
	RegisterHealthRoutes()

	// Bind API documentation endpoints.
	RegisterSwaggerRoutes()
}
