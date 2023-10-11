package es_utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

type SearchResult struct {
	Response *esapi.Response
	Error    error
}

func Search(searchTerm string) (map[string]interface{}, error) {
	// Initialize Elasticsearch client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	// Define the fields to search and the prefix query
	fieldsToSearch := []string{"PatentTitle", "Authors", "Assignee", "DesignClass"}
	prefixQuery := "Art" // Prefix for the search

	// Create a bool query with should clauses for each field
	shouldClauses := make([]map[string]interface{}, len(fieldsToSearch))
	for i, field := range fieldsToSearch {
		wildcardQueryForField := map[string]interface{}{
			"wildcard": map[string]interface{}{
				field: "*" + strings.ToLower(prefixQuery) + "*", // Case-insensitive wildcard query
			},
		}
		shouldClauses[i] = map[string]interface{}{"bool": wildcardQueryForField}
	}

	// ...

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"ApplicationDate": map[string]interface{}{
								"query":     prefixQuery,
								"fuzziness": "AUTO",
							},
						},
					},
					{
						"match": map[string]interface{}{
							"Assignee": map[string]interface{}{
								"query":     prefixQuery,
								"fuzziness": "AUTO",
							},
						},
					},
					{
						"match": map[string]interface{}{
							"Authors": map[string]interface{}{
								"query":     prefixQuery,
								"fuzziness": "AUTO",
							},
						},
					},
					{
						"match": map[string]interface{}{
							"DesignClass": map[string]interface{}{
								"query":     prefixQuery,
								"fuzziness": "AUTO",
							},
						},
					},
					{
						"match": map[string]interface{}{
							"IssueDate": map[string]interface{}{
								"query":     prefixQuery,
								"fuzziness": "AUTO",
							},
						},
					},
					{
						"match": map[string]interface{}{
							"PatentNumber": map[string]interface{}{
								"query":     prefixQuery,
								"fuzziness": "AUTO",
							},
						},
					},
					{
						"match": map[string]interface{}{
							"PatentTitle": map[string]interface{}{
								"query":     prefixQuery,
								"fuzziness": "AUTO",
							},
						},
					},
				},
			},
		},
	}

	// ...

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		log.Fatalf("Error encoding the search query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("design_patent"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error performing the search request: %v", err)
	}
	defer res.Body.Close()

	// Read the response body into a byte slice
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body: %v", err)
	}

	// Check for successful status code from Elasticsearch
	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch request failed with status code: %d", res.StatusCode)
	}

	// Unmarshal the response body into a map
	var responseMap map[string]interface{}
	if err := json.Unmarshal(responseBody, &responseMap); err != nil {
		return nil, fmt.Errorf("error decoding Elasticsearch response: %v", err)
	}

	return responseMap, nil
}

func SearchWithPagination(searchTerm string, from, size int) (map[string]interface{}, error) {
	// Initialize Elasticsearch client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	// Define the fields to search and the prefix query
	fieldsToSearch := []string{"PatentTitle", "Authors", "Assignee", "DesignClass"}
	prefixQuery := strings.ToLower(searchTerm) // Adjust prefixQuery to use the provided search term

	// Create a bool query with should clauses for each field
	shouldClauses := make([]map[string]interface{}, len(fieldsToSearch))
	for i, field := range fieldsToSearch {
		wildcardQueryForField := map[string]interface{}{
			"wildcard": map[string]interface{}{
				field: "*" + prefixQuery + "*", // Case-insensitive wildcard query
			},
		}
		shouldClauses[i] = map[string]interface{}{"bool": wildcardQueryForField}
	}

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": shouldClauses,
			},
		},
	}

	// Encode the query
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, fmt.Errorf("error encoding the search query: %v", err)
	}

	// Perform the search request with pagination
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("design_patent"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithFrom(from),
		es.Search.WithSize(size),
	)
	if err != nil {
		return nil, fmt.Errorf("error performing the search request: %v", err)
	}
	defer res.Body.Close()

	// Read the response body into a byte slice
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body: %v", err)
	}

	// Check for successful status code from Elasticsearch
	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch request failed with status code: %d", res.StatusCode)
	}

	// Unmarshal the response body into a map
	var responseMap map[string]interface{}
	if err := json.Unmarshal(responseBody, &responseMap); err != nil {
		return nil, fmt.Errorf("error decoding Elasticsearch response: %v", err)
	}

	return responseMap, nil
}
