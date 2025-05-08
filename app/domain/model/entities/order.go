package entities

import (
	"time"
)

type Order struct {
	PurchaseId  string
	BookingId   string
	Status      string
	TotalAmount int64
	Description string
	CreatedAt   time.Time
}

func ToMapOrders(items []map[string]any) map[string][]Order {
	orders := make(map[string][]Order)
	for _, item := range items {
		order := ToOrder(item)
		orders[order.BookingId] = append(orders[order.BookingId], order)
	}
	return orders
}

func ToOrder(item map[string]any) Order {
	return Order{
		PurchaseId:  item["purchase_id"].(string),
		BookingId:   item["booking_id"].(string),
		Status:      item["status"].(string),
		TotalAmount: item["total_amount"].(int64),
		Description: item["description"].(string),
		CreatedAt:   item["created_at"].(time.Time),
	}
}
