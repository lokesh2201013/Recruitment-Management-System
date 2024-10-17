package main

import (
	"log"
    "github.com/lokesh2201013/config"
    "github.com/lokesh2201013/routes"
    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)
func init() {
	// Load environment variables from the .env file
	/*if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading env file")
	}*/

	// Connect to the database
	config.ConnectDB()
}

func main() {

	sqlDb, err := config.DB.DB()
    if err != nil {
        panic("Error in SQL connection")
    }
    defer sqlDb.Close()

	
    app := fiber.New()
    
	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

	app.Use(logger.New())

    routes.AuthRoutes(app)

	if err := app.Listen(":8080"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}