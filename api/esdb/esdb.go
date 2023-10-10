package esdb

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var esClient *elasticsearch.Client

func InitElasticsearch() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
		return nil, err
	}

	esClient = client
	return client, nil
}

func GetESClient() *elasticsearch.Client {
	return esClient
}
