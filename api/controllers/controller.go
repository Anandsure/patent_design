package controllers

import (
	"fmt"

	"github.com/Anandsure/patent_design/api/esdb"
	"github.com/gofiber/fiber/v2"
)

func SearchHandler(c *fiber.Ctx) error {
	searchTerm := c.Query("q") // Get the search term from the URL parameter 'q'

	// Perform the Elasticsearch query
	results, err := esdb.Search(searchTerm)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to perform Elasticsearch query: %v", err),
		})
	}

	return c.JSON(fiber.Map{
		"results": results,
	})
}
