package elasticsearch

import (
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func CreateDocument(id string, body interface{}, indexNames ...string) (res *esapi.Response, err error) {
	b, err := json.Marshal(body)
	if err != nil {
		return
	}
	payload := string(b)

	index := createIndex(indexNames...)
	return Client.Create(index, id, strings.NewReader(payload))
}
