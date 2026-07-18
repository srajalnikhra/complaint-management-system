package routes

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/controllers"
	"github.com/srajalnikhra/complaint-management-system/internal/middleware"
)

// RegisterUserRoutes declares user management, authentication, and password reset endpoints.
func RegisterUserRoutes() {

	// Create controllers dependencies.
	userController := controllers.NewUserController()
	loginController := controllers.NewLoginController()

	// Register public signup router endpoint.
	http.HandleFunc("/register", userController.Register)

	// Register authentication and credentials validation endpoint wrapped in the RateLimiter.
	http.Handle(
		"/login",
		middleware.RateLimitMiddleware(
			http.HandlerFunc(loginController.Login),
		),
	)

	// Register OTP password recovery pathways protected against abuse.
	http.Handle(
		"/forgot-password",
		middleware.RateLimitMiddleware(
			http.HandlerFunc(loginController.ForgotPassword),
		),
	)

	http.Handle(
		"/verify-otp",
		middleware.RateLimitMiddleware(
			http.HandlerFunc(loginController.VerifyOTP),
		),
	)

	http.Handle(
		"/reset-password",
		middleware.RateLimitMiddleware(
			http.HandlerFunc(loginController.ResetPassword),
		),
	)
}
