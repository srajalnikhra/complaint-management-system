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

	_ "github.com/srajalnikhra/complaint-management-system/docs"
	"github.com/srajalnikhra/complaint-management-system/internal/config"
	"github.com/srajalnikhra/complaint-management-system/internal/database"
	"github.com/srajalnikhra/complaint-management-system/internal/middleware"
	"github.com/srajalnikhra/complaint-management-system/internal/routes"
)

func main() {

	config.Initialize()

	appConfig := config.LoadAppConfig()
	dbConfig := config.LoadDBConfig()

	database.Initialize(dbConfig)

	routes.RegisterRoutes()

	go middleware.StartRateLimiterCleanup()

	handler := middleware.LoggingMiddleware(
		middleware.CORSMiddleware(
			http.DefaultServeMux,
		),
	)

	server := &http.Server{
		Addr:    ":" + appConfig.Port,
		Handler: handler,
	}

	go func() {

		log.Printf("%s started on port %s", appConfig.Name, appConfig.Port)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-stop

	log.Println("Shutting down server...")

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
