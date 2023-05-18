package elastic

import (
	elastic "airbnb-auth-be/internal/pkg/elasticsearch"
	"airbnb-auth-be/internal/pkg/log"
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const Instance = "Elastic Middleware"

var indexName = []string{"request", "api"}

func LogRequestToElastic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Create a log writter
		logWriter := &GinLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = logWriter

		// Create request log data
		startTime := time.Now()
		requestBody, _ := ioutil.ReadAll(ctx.Request.Body)
		reader := ioutil.NopCloser(bytes.NewBuffer(requestBody))
		ctx.Request.Body = reader
		request := Request{
			Time:      startTime,
			Method:    ctx.Request.Method,
			Uri:       ctx.Request.RequestURI,
			Proto:     ctx.Request.Proto,
			UserAgent: ctx.Request.UserAgent(),
			Referer:   ctx.Request.Referer(),
			Body:      string(requestBody),
			Ip:        ctx.ClientIP(),
		}

		// Processing the request
		ctx.Next()

		// Create response log data
		endTime := time.Now()
		response := Response{
			Time:       endTime,
			StatusCode: ctx.Writer.Status(),
			Body:       logWriter.body.String(),
		}

		body := Log{
			ProcessTime: endTime.Sub(startTime),
			Request:     request,
			Response:    response,
		}

		go func() {
			id, err := gonanoid.New()
			if err != nil {
				return
			}

			res, err := elastic.CreateDocument(id, body, indexName...)
			if err != nil || res.IsError() {
				msg := fmt.Sprintf("error indexing for document id=%s", id)
				log.Error(Instance, msg, fmt.Errorf("%v", res))
				return
			}
		}()
	}
}

func CreateIndex() {
	res, err := elastic.IsIndexExist(indexName...)
	if err != nil {
		log.Fatal(Instance, "failed to check index", fmt.Errorf("%v", res))
	} else if res.StatusCode == 404 {
		res, err = elastic.CreateIndex(esMapping, indexName...)
		if err != nil || res.IsError() {
			log.Fatal(Instance, "failed to create index", fmt.Errorf("%v", res))
		}
	}
}
