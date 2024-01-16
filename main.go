package main

import (
	"log"

	"github.com/Pebody-h/Fake-Shop-Api/database"
	"github.com/Pebody-h/Fake-Shop-Api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
