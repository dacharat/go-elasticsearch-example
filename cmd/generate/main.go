package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/dacharat/go-elasticsearch-example/internal/mock"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
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

	loadData(es, 1000) // http://localhost:9200/order/_search
}

func loadData(es *elasticsearch.Client, n int) {
	orders := mock.GenerateOrders(n)

	for _, data := range orders {
		orderID := data.OrderID
		jsonString, _ := json.Marshal(data)
		request := esapi.IndexRequest{Index: "order", DocumentID: orderID, Body: strings.NewReader(string(jsonString))}
		request.Do(context.Background(), es)
	}
	fmt.Println(len(orders), " orders read")
}
