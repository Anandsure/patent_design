package es_utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/Anandsure/patent_design/api/db"

	"github.com/elastic/go-elasticsearch/esapi"
)

type SearchResult struct {
	Response *esapi.Response
	Error    error
}

func Search(searchTerm string) (map[string]interface{}, error) {
	es := db.GetESClient()

	// Define the fields to search
	fieldsToSearch := []string{"PatentTitle", "Authors", "Assignee", "DesignClass", "ApplicationDate", "IssueDate", "PatentNumber"}

	// Construct the query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query":     fmt.Sprintf("*%s*", strings.ToLower(searchTerm)),
				"fields":    fieldsToSearch,
				"fuzziness": "AUTO",
			},
		},
	}

	// Encode the query
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("error encoding the search query: %v", err)
	}

	// Perform the search request
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(db.ES_INDEX_NAME),
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
