package entities

import (
	"time"
)

type Booking struct {
	BookingId       string
	ClientId        int64
	Status          string
	OriginPort      string
	DestinationPort string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func ToMapBookings(items []map[string]any) map[string]Booking {
	bookings := make(map[string]Booking)
	for _, item := range items {
		bookings[item["booking_id"].(string)] = ToBooking(item)
	}
	return bookings
}

func ToBooking(item map[string]any) Booking {
	return Booking{
		BookingId:       item["booking_id"].(string),
		ClientId:        item["client_id"].(int64),
		Status:          item["status"].(string),
		OriginPort:      item["origin_port"].(string),
		DestinationPort: item["destination_port"].(string),
		CreatedAt:       item["created_at"].(time.Time),
		UpdatedAt:       item["updated_at"].(time.Time),
	}
}
