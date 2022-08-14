package mock

import (
	"fmt"
	"time"

	"github.com/dacharat/go-elasticsearch-example/internal/model"
	"github.com/dacharat/go-elasticsearch-example/internal/util/random"
	"github.com/google/uuid"
)

var buz = []model.Address{
	{
		Name:        "Buz-1-one",
		Phone:       "01111111111",
		AddressName: "Thailand",
	},
	{
		Name:        "Buz-2-two",
		Phone:       "0222222222",
		AddressName: "Germany",
	},
	{
		Name:        "Buz-3-three",
		Phone:       "033333333",
		AddressName: "Swiss",
	},
	{
		Name:        "Buz-4-four",
		Phone:       "0444444444",
		AddressName: "Singapore",
	},
	{
		Name:        "Buz-5-five",
		Phone:       "0555555555",
		AddressName: "Barcelona",
	},
}

var cus = []model.Address{
	{
		Name:        "Cus-1-lewi",
		Phone:       "0811111111",
		AddressName: "Bangkok",
	},
	{
		Name:        "Cus-2-rafinha",
		Phone:       "0822222222",
		AddressName: "Berlin",
	},
	{
		Name:        "Cus-3-dem",
		Phone:       "083333333",
		AddressName: "Zermatt",
	},
	{
		Name:        "Cus-4-fdy",
		Phone:       "0844444444",
		AddressName: "Singapore",
	},
	{
		Name:        "Cus-5-messi",
		Phone:       "0855555555",
		AddressName: "Camp nou",
	},
}

func GenerateOrders(n int) []model.Order {
	orders := make([]model.Order, n)
	for i := range orders {
		orders[i] = model.Order{
			OrderID: fmt.Sprintf("Order-%s", uuid.New()),
			Trips: []model.Address{
				random.Random(buz),
				random.Random(cus),
			},
			Status:     random.Random([]model.OrderStatus{model.StatusCanceled, model.StatusCompleted, model.StatusFailed}),
			StatusTime: time.Now(),
		}
	}
	return orders
}
