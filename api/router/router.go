package router

import (
	"github.com/Anandsure/patent_design/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app fiber.Router) {
	// Add the /search route
	app.Get("/search", controllers.SearchHandler)
	// Add the /query route
	app.Get("/query", controllers.QueryHandler)
}
