// @title Complaint Management System API
// @version 1.0
// @description Complaint Management System Backend API
// @host localhost:8080
// @BasePath /
//
// @tag.name Health
// @tag.description Health check endpoints
//
// @tag.name Authentication
// @tag.description Authentication and account recovery
//
// @tag.name Complaints
// @tag.description Complaint management endpoints
//
// @tag.name Admin
// @tag.description Administrator operations
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/srajalnikhra/complaint-management-system/docs"
	_ "github.com/srajalnikhra/complaint-management-system/docs"
	"github.com/srajalnikhra/complaint-management-system/internal/config"
	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/middleware"
	"github.com/srajalnikhra/complaint-management-system/internal/routes"
)

// main starts the HTTP server, connects to the database, and sets up
// routes and middleware. It also listens for termination signals to shutdown
// the server and close database connections cleanly.
func main() {

	// Load the application configuration settings.
	config.Initialize()

	appConfig := config.LoadAppConfig()
	dbConfig := config.LoadDBConfig()
	
	// Application changes the Swagger host at runtime.
	if os.Getenv("APP_ENV") == "production" {
		docs.SwaggerInfo.Host = "complaint-management-system-lciv.onrender.com"
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.Schemes = []string{"http"}
	}

	// Connect to PostgreSQL database and run table migrations.
	database.Initialize(dbConfig)

	// Register all API routes.
	routes.RegisterRoutes()

	// Start a background worker to clean up expired rate limit records.
	go middleware.StartRateLimiterCleanup()

	// Wrap the multiplexer handlers in CORS and logging middleware.
	handler := middleware.LoggingMiddleware(
		middleware.CORSMiddleware(
			http.DefaultServeMux,
		),
	)

	server := &http.Server{
		Addr:    ":" + appConfig.Port,
		Handler: handler,
	}

	// Start the HTTP server in a separate goroutine so it doesn't block the main flow.
	go func() {

		log.Printf("%s started on port %s", appConfig.Name, appConfig.Port)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Listen for system signals to support graceful shutdown.
	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		os.Interrupt,
		syscall.SIGTERM,
	)

	// Wait until we receive a termination signal.
	<-stop

	log.Println("Shutting down server...")

	// Gracefully shut down the server, allowing active requests to complete within 10 seconds.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	database.DB.Close()

	log.Println("Database connection closed")

	log.Println("Server stopped gracefully")
}
