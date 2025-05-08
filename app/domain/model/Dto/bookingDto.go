package Dto

import "prueba-tecnica-nauta/app/domain/model/entities"

type BookingDto struct {
	BookingId       string         `json:"booking_id"`
	ClientId        int64          `json:"client_id"`
	Status          string         `json:"status"`
	OriginPort      string         `json:"origin_port"`
	DestinationPort string         `json:"destination_port"`
	Containers      []ContainerDto `json:"containers"`
	Orders          []OrderDto     `json:"orders"`
}

func FromBooking(booking entities.Booking, containers []entities.Container, orders []entities.Order, invoices map[string][]entities.Invoice) BookingDto {
	bookingContainers := make([]ContainerDto, len(containers))
	for i, container := range containers {
		if container.BookingId == booking.BookingId {
			bookingContainers[i] = FromContainer(container)
		}
	}

	bookingOrders := make([]OrderDto, len(orders))
	for i, order := range orders {
		if order.BookingId == booking.BookingId {
			bookingOrders[i] = FromOrder(order, invoices[order.PurchaseId])
		}
	}

	return BookingDto{
		BookingId:       booking.BookingId,
		Status:          booking.Status,
		ClientId:        booking.ClientId,
		OriginPort:      booking.OriginPort,
		DestinationPort: booking.DestinationPort,
		Containers:      bookingContainers,
		Orders:          bookingOrders,
	}
}

func (b *BookingDto) ToBooking(clientId int64) entities.Booking {
	return entities.Booking{
		BookingId:       b.BookingId,
		ClientId:        b.ClientId,
		Status:          b.Status,
		OriginPort:      b.OriginPort,
		DestinationPort: b.DestinationPort,
	}
}
