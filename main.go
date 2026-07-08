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

	config.LoadEnv()

	appConfig := config.LoadAppConfig()

	dbConfig := config.LoadDBConfig()

	database.ConnectDB(dbConfig)

	database.RunMigrations()

	routes.RegisterUserRoutes()
	routes.RegisterProtectedRoutes()
	routes.RegisterComplaintRoutes()
	routes.RegisterAdminRoutes()
	routes.RegisterHealthRoutes()

	log.Printf("%s started on port %s", appConfig.Name, appConfig.Port)

	handler := middleware.LoggingMiddleware(
		middleware.CORSMiddleware(
			http.DefaultServeMux,
		),
	)

	log.Fatal(http.ListenAndServe("localhost:"+appConfig.Port, handler))
}
