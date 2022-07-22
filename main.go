package main

import (
	"log"
	"os"

	"github.com/CamielaTeam/go-claridad-api/controllers/buildings"
	"github.com/CamielaTeam/go-claridad-api/db"
	"github.com/CamielaTeam/go-claridad-api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()

	dbErr := db.Open()

	migrationsErr := db.Automigrate()

	if envErr != nil {
		log.Fatal("Error loading environment file, ", envErr)
	}

	if dbErr != nil {
		log.Fatal("Error connecting to database, ", dbErr)
	}

	if migrationsErr != nil {
		log.Fatal("Error executing migrations, ", migrationsErr)
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(middlewares.RequestIDMiddleware())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r.Use(cors.New(corsConfig))

	buildings.RegisterRoutes(r)

	r.Run()
}
