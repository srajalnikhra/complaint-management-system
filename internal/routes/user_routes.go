package routes

import (
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/controllers"
)

func RegisterUserRoutes() {

	userController := controllers.NewUserController()
	loginController := controllers.NewLoginController()

	http.HandleFunc("/register", userController.Register)
	http.HandleFunc("/login", loginController.Login)
	http.HandleFunc("/forgot-password", loginController.ForgotPassword)
	http.HandleFunc("/verify-otp", loginController.VerifyOTP)
	http.HandleFunc("/reset-password", loginController.ResetPassword)
}