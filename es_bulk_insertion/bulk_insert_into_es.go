package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func main() {
	// Read data from JSON file
	filePath := "../file_extraction/json_extraction/combined_patent_data.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %s", err)
	}

	// Initialize Elasticsearch client
	client, err := initElasticsearch()
	if err != nil {
		log.Fatalf("Error initializing Elasticsearch: %v", err)
	}

	// Parse JSON data
	var docs []map[string]interface{}
	if err := json.Unmarshal(data, &docs); err != nil {
		log.Fatalf("Error parsing JSON data: %v", err)
	}

	// Bulk insert the documents
	if err := bulkInsert(client, docs); err != nil {
		log.Fatalf("Error performing bulk insert: %v", err)
	}

	fmt.Println("Bulk insert completed successfully.")
}

func initElasticsearch() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func bulkInsert(client *elasticsearch.Client, docs []map[string]interface{}) error {
	var bulkBody []byte

	for _, doc := range docs {
		metaData := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": "patent_data",
			},
		}

		metaDataBytes, err := json.Marshal(metaData)
		if err != nil {
			return err
		}

		docBytes, err := json.Marshal(doc)
		if err != nil {
			return err
		}

		bulkBody = append(bulkBody, metaDataBytes...)
		bulkBody = append(bulkBody, []byte("\n")...)
		bulkBody = append(bulkBody, docBytes...)
		bulkBody = append(bulkBody, []byte("\n")...)
	}

	req := esapi.BulkRequest{
		Body:    bytes.NewReader(bulkBody),
		Index:   "patent_data",
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("bulk request failed with status: %s", res.Status())
	}

	return nil
}
