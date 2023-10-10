package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	// Initialize the Elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{viper.GetString("ES_HOST")},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	// Read data from JSON file
	filePath := "../file_extraction/json_extraction/combined_patent_data.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %s", err)
	}

	// Unmarshal JSON data
	var jsonData []map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		log.Fatalf("Error unmarshaling JSON: %s", err)
	}

	// Bulk index request
	var bulkRequestBody bytes.Buffer

	for _, doc := range jsonData {
		// Exclude Description and ReferencesCited fields
		delete(doc, "Description")
		delete(doc, "ReferencesCited")

		metaData := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": "design_patent", // Change the index name here
			},
		}

		if err := json.NewEncoder(&bulkRequestBody).Encode(metaData); err != nil {
			log.Fatalf("Error encoding metadata: %s", err)
		}

		if err := json.NewEncoder(&bulkRequestBody).Encode(doc); err != nil {
			log.Fatalf("Error encoding document: %s", err)
		}
	}

	// Perform the bulk request
	res, err := es.Bulk(bytes.NewReader(bulkRequestBody.Bytes()), es.Bulk.WithContext(context.Background()))
	if err != nil {
		log.Fatalf("Error performing bulk insert: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var buf bytes.Buffer
		io.Copy(&buf, res.Body)
		log.Fatalf("Error response: %s", buf.String())
	}

	log.Println("Bulk insertion completed successfully.")
}
