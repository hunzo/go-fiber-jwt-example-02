package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hunzo/go-fiber-jwt-example-02/handlers"
	"github.com/hunzo/go-fiber-jwt-example-02/validate"
)

func SetupRouters(r *fiber.App) {
	public := r.Group("/")
	public.Get("/login", handlers.Login)

	private := r.Group("/api", validate.AuthedRequired())
	private.Get("/profile", handlers.Profile)

}
