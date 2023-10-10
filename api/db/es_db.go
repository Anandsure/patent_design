package db

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/spf13/viper"
	"log"
)

var esClient *elasticsearch.Client

var ES_INDEX_NAME string

func InitElasticsearch() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{viper.GetString("ES_HOST")},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
		return nil, err
	}

	ES_INDEX_NAME = viper.GetString("ES_INDEX_NAME")

	esClient = client
	return client, nil
}

func GetESClient() *elasticsearch.Client {
	return esClient
}
