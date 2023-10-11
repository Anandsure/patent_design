package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Anandsure/patent_design/api/db"
	"github.com/Anandsure/patent_design/api/es_utils"
	"github.com/gofiber/fiber/v2"
)

type SearchResponse struct {
	Results map[string]interface{} `json:"results"`
}

func SearchHandler(c *fiber.Ctx) error {
	// Get the search term and pagination parameters from the URL query
	searchTerm := c.Query("q")                      // Get the search term from the URL parameter 'q'
	from, err := strconv.Atoi(c.Query("from", "0")) // Get the 'from' parameter for pagination (default: 0)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid 'from' parameter. Please provide a valid integer.",
		})
	}
	size, err := strconv.Atoi(c.Query("size", "10")) // Get the 'size' parameter for pagination (default: 10)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid 'size' parameter. Please provide a valid integer.",
		})
	}

	// Perform the Elasticsearch query with pagination
	results, err := es_utils.SearchWithPagination(searchTerm, from, size)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to perform Elasticsearch query: %v", err),
		})
	}
	hitsInterface, hitsExist := results["hits"]
	if !hitsExist {
		return c.Status(500).JSON(fiber.Map{
			"error": "Elasticsearch query results do not contain 'hits'",
		})
	}

	// Convert the hits to a JSON-serializable format
	hitsJSON, err := json.Marshal(hitsInterface)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to marshal hits to JSON: %v", err),
		})
	}

	// Respond with the hits as JSON
	return c.Send(hitsJSON)

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
