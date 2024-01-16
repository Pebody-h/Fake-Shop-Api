package router

import (
	"github.com/Pebody-h/Fake-Shop-Api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// ** Rutas
	app.Get("/test", handlers.ConnectionTest)
	app.Get("/api/products", handlers.GetProducts)
	app.Get("/api/products/:id", handlers.GetProductById)
	app.Post("/api/product", handlers.PostProduct)
	app.Post("/fill/database", handlers.PostProducts)
}
