package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/Anandsure/patent_design/api/db"
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
		"results": results["hits"],
	})
}
func QueryHandler(c *fiber.Ctx) error {
	searchTerm := c.Query("patentnumber")
	resultJSON, err := db.GetPatentJSONByNumber(searchTerm)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to perform PostgreSQL query: %v", err),
		})
	}

	// Check if patent is found
	if resultJSON == "" {
		return c.Status(404).JSON(fiber.Map{
			"error": fmt.Sprintf("No patent found with patent number: %s", searchTerm),
		})
	}

	// Parse the JSON string to a map
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(resultJSON), &result); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to parse JSON response: %v", err),
		})
	}

	return c.JSON(result)
}
