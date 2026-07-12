package routes

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/controllers"
	"github.com/srajalnikhra/complaint-management-system/internal/middleware"
)

func RegisterUserRoutes() {

	userController := controllers.NewUserController()
	loginController := controllers.NewLoginController()

	http.HandleFunc("/register", userController.Register)
	http.Handle(
		"/login",
		middleware.RateLimitMiddleware(
			http.HandlerFunc(loginController.Login),
		),
	)

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
