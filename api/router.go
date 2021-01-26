package api

import (
	"github.com/cjtim/golang-fiber-mongodb/api/controllers"

	"github.com/gofiber/fiber/v2"
)

// Route for all api request
func Route(r *fiber.App) {
	r.Get("/", controllers.HomeController)
	r.Get("/ping", controllers.PingController)
	r.Post("/post", controllers.PostController)
}
