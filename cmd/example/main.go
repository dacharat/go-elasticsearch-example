package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/dacharat/go-elasticsearch-example/internal/data"
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	log.Println(elasticsearch.Version)

	// res, err := es.Ping()
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error ping response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)

	// data.LoadData(es) // http://localhost:9200/stsc/_search

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("0) Exit")
		fmt.Println("1) Load spacecraft")
		fmt.Println("2) Get spacecraft")
		fmt.Println("3) Search spacecraft by key and value")
		fmt.Println("4) Search spacecraft by key and prefix")
		option := data.ReadText(reader, "Enter option")
		if option == "0" {
			data.Exit()
		} else if option == "1" {
			data.LoadData(es)
		} else if option == "2" {
			data.Get(es, reader)
		} else if option == "3" {
			data.Search(es, reader, "match")
		} else if option == "4" {
			data.Search(es, reader, "prefix")
		} else {
			fmt.Println("Invalid option")
		}
	}
}
