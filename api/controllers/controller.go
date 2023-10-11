package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/Anandsure/patent_design/api/db"
	"github.com/Anandsure/patent_design/api/es_utils"
	"github.com/gofiber/fiber/v2"
)

type SearchResponse struct {
	Results map[string]interface{} `json:"results"`
}

func SearchHandler(c *fiber.Ctx) error {
	searchTerm := c.Query("q") // Get the search term from the URL parameter 'q'

	// Perform the Elasticsearch query
	fmt.Println(searchTerm)
	results, err := es_utils.Search(searchTerm)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to perform Elasticsearch query: %v", err),
		})
	}
	fmt.Printf("Search poppers: %+v\n", results)
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
