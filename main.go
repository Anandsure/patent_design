package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"

	"github.com/Anandsure/patent_design/api/cache"
	"github.com/Anandsure/patent_design/api/controllers"
	"github.com/Anandsure/patent_design/api/db"
	"github.com/Anandsure/patent_design/api/migrations"
	"github.com/Anandsure/patent_design/api/utils"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("Running fine bro")
}

func main() {
	// Set global configuration
	utils.ImportEnv()

	// Init redis
	cache.GetRedis()

	// Init Validators
	utils.InitValidators()

	// Create Fiber
	app := fiber.New(fiber.Config{})

	app.Get("/", healthCheck)
	app.Get("/health", healthCheck)
	// Add the /search route
	app.Get("/search", controllers.SearchHandler)
	// Add the /query route
	app.Get("/query", controllers.QueryHandler)

	app.Use(logger.New(logger.Config{Next: func(c *fiber.Ctx) bool {
		return strings.HasPrefix(c.Path(), "api")
	}}))

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowHeaders: "*",
	}))

	//Connect and migrate the db
	if viper.GetBool("MIGRATE") {
		migrations.Migrate()
	}

	// Initialize DB
	db.InitServices()

	// Get Port
	port := utils.GetPort()

	// Start Fiber
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	// r := router.SetupRouter()
	// r.Run(":8080")

}
