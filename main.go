package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hunzo/go-fiber-jwt-example-02/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	routes.SetupRouters(app)

	app.Listen(":8080")
}
