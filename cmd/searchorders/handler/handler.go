package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	es *elasticsearch.Client
}

func NewHandler(es *elasticsearch.Client) Handler {
	return Handler{es: es}
}

func (h Handler) SearchOrderHandler(c *gin.Context) {
	q := c.Query("q")
	fromStr := c.Query("from")
	sizeStr := c.Query("size")

	var (
		from int
		size = 20
	)

	if fromStr != "" {
		n, err := strconv.Atoi(fromStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		from = n
	}

	if sizeStr != "" {
		n, err := strconv.Atoi(sizeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		size = n
	}

	var buffer bytes.Buffer
	query := map[string]interface{}{
		"from": from,
		"size": size,
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  q,
				"fields": []string{"status", "trips.name", "trips.phone", "trips.addressName"},
			},
		},
	}
	json.NewEncoder(&buffer).Encode(query)
	response, _ := h.es.Search(h.es.Search.WithIndex("order"), h.es.Search.WithBody(&buffer))
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	c.JSON(http.StatusOK, result)
}
