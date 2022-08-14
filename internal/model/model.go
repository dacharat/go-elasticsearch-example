package model

import "time"

type OrderStatus string

const (
	StatusCompleted OrderStatus = "COMPLETED"
	StatusCanceled  OrderStatus = "CANCELED"
	StatusFailed    OrderStatus = "FAILED"
)

type Order struct {
	OrderID    string      `json:"orderId"`
	Trips      []Address   `json:"trips"`
	Status     OrderStatus `json:"status"`
	StatusTime time.Time   `json:"statusTime"`
}

type Address struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	AddressName string `json:"addressName"`
}
