package main

import (
	"log"
	"net/http"

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

	log.Printf("%s started on port %s", appConfig.Name, appConfig.Port)

	log.Fatal(http.ListenAndServe("localhost:"+appConfig.Port, handler))
}
