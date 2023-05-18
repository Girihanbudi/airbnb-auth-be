package elasticsearch

import (
	"airbnb-auth-be/internal/pkg/env"
	"airbnb-auth-be/internal/pkg/log"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
)

const Instance string = "Elasticsearch Client"

// global elastic search client declaration
var Client *elasticsearch.Client

func InitElasticSearch() {

	config := elasticsearch.Config{
		Addresses: env.CONFIG.Elastic.Addresses,
		Username:  env.CONFIG.Elastic.Username,
		Password:  env.CONFIG.Elastic.Password,
	}

	client, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatal(Instance, "connection error", err)
	}

	_, err = client.Info()
	if err != nil {
		log.Fatal(Instance, "failed to get client info", err)
	}

	// log.Event(Instance, fmt.Sprintf("connected to %v, client info: %v", env.CONFIG.Elastic.Addresses, res))
	log.Event(Instance, fmt.Sprintf("connected to %v", env.CONFIG.Elastic.Addresses))
	Client = client
}
