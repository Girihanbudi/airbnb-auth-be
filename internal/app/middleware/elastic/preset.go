package elastic

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

type Log struct {
	ProcessTime time.Duration `json:"process_time"`
	Request     Request       `json:"request"`
	Response    Response      `json:"response"`
}

type Request struct {
	Time      time.Time `json:"time"`
	Method    string    `json:"method"`
	Uri       string    `json:"uri"`
	Proto     string    `json:"proto"`
	UserAgent string    `json:"user_agent"`
	Referer   string    `json:"referer"`
	Body      string    `json:"body"`
	Ip        string    `json:"ip"`
}

type Response struct {
	Time       time.Time `json:"time"`
	StatusCode int       `json:"status_code"`
	Body       string    `json:"body"`
}

type GinLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

const esMapping = `
	{
		"settings":{
			"number_of_shards": 1,
			"number_of_replicas": 1
		}
	}`
