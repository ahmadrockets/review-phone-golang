package main

import (
	"final-project-sanber/config"
	"final-project-sanber/docs"
	"final-project-sanber/routes"
	"final-project-sanber/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// for load godotenv
	// for env
	environment := utils.Getenv("ENVIRONMENT", "development")
	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger API For Phone Review"
	docs.SwaggerInfo.Description = "This is a simple server for phone by Ahmad Fahrudin."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// database connection
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Seeder open when first time install app
	/* for _, rseed := range seed.All() {
		if err := rseed.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", rseed.Name, err)
		}
	} */

	// router
	r := routes.SetupRouter(db)
	// just remove port 8080
	r.Run()
}
