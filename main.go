package main

import (
	"fmt"
	"strings"

	"github.com/Anandsure/patent_design/api/router"
	"github.com/Anandsure/patent_design/bulk_insertion"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"

	"github.com/Anandsure/patent_design/api/db"
	"github.com/Anandsure/patent_design/api/migrations"
	"github.com/Anandsure/patent_design/api/utils"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("Running fine bro")
}

func startServer() {
	// Set global configuration
	utils.ImportEnv()

	// Init Validators
	utils.InitValidators()

	// Create Fiber
	app := fiber.New(fiber.Config{})

	app.Get("/", healthCheck)
	app.Get("/health", healthCheck)

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

	// Initialize DB and ES
	db.InitServices()

	// Get Port
	port := utils.GetPort()

	router.MountRoutes(app)

	// Start Fiber
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}

func main() {
	if viper.GetBool("START_SERVER") {
		startServer()
	} else {
		bulk_insertion.StartBulk()
	}
}
