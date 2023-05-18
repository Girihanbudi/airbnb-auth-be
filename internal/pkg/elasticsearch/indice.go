package elasticsearch

import (
	"airbnb-auth-be/internal/pkg/env"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func createIndex(names ...string) string {
	separator := env.CONFIG.Elastic.Separator
	if separator == "" {
		separator = "_"
	}

	return strings.Join(names, separator)
}

func CreateIndex(mapping string, indexNames ...string) (*esapi.Response, error) {
	index := createIndex(indexNames...)
	return Client.Indices.Create(index, Client.Indices.Create.WithBody(strings.NewReader(mapping)))
}

func IsIndexExist(indexNames ...string) (*esapi.Response, error) {
	index := createIndex(indexNames...)
	return Client.Indices.Exists([]string{index})
}
