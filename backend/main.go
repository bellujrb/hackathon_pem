
package main

import (
	"perna/db"
	"perna/middleware"
	"log"
	"os"

	"github.com/swaggo/swag/example/basic/docs"
	"gorm.io/gorm"
)

// @title           perna
// @version         1.0
// @description     This is a server for app.

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	dbInstance := db.ConnectDatabaseGorm()

	r := middleware.SetupRouter(dbInstance)

	migrate(dbInstance)

	host := os.Getenv("API_HOST")
	if host == "" {
		host = "localhost:8080"
	}
	docs.SwaggerInfo.Host = host

	r.Run(":8080")
}

func migrate(dbInstance *gorm.DB) {
	if err := db.Migrate(dbInstance); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
    