package main

import (
	"log"
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/config"
	"github.com/srajalnikhra/complaint-management-system/internal/database"
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

	log.Printf("%s started on port %s", appConfig.Name, appConfig.Port)

	log.Fatal(http.ListenAndServe(":"+appConfig.Port, nil))
}
