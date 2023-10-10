package controllers

import (
	"fmt"

	"github.com/Anandsure/patent_design/api/db"
	"github.com/Anandsure/patent_design/api/es_utils"
	"github.com/gofiber/fiber/v2"
)

func SearchHandler(c *fiber.Ctx) error {
	searchTerm := c.Query("q") // Get the search term from the URL parameter 'q'

	// Perform the Elasticsearch query
	results, err := es_utils.Search(searchTerm)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to perform Elasticsearch query: %v", err),
		})
	}

	return c.JSON(fiber.Map{
		"results": results["hits"],
	})
}
func QueryHandler(c *fiber.Ctx) error {
	searchTerm := c.Query("patent_number", "")

	result, err := db.PatentSvc.GetPatentJSONByNumber(searchTerm)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to perform PostgreSQL query: %v", err),
		})
	}

	return c.JSON(result)
}
