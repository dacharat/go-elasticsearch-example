package main

import (
	"log"

	"github.com/dacharat/go-elasticsearch-example/cmd/searchorders/handler"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	log.Println(elasticsearch.Version)

	res, err := es.Ping()
	if err != nil {
		log.Fatalf("Error ping response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)

	h := handler.NewHandler(es)

	router := gin.New()

	router.GET("search", h.SearchOrderHandler)

	router.Run()
}
