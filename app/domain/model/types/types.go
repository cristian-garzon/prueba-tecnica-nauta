package types

import (
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/entities"
)

// sqlRepository
type FindAllQuery func(query string, args ...any) ([]map[string]any, error)
type FindOneQuery func(query string, args ...any) (map[string]any, error)
type ExecuteQuery func(query string, args ...any) (map[string]any, error)

// cacheRepository
type MapBookings func() (map[string]entities.Booking, error)
type MapContainers func() (map[string][]entities.Container, error)
type MapOrders func() (map[string][]entities.Order, error)
type MapInvoices func() (map[string][]entities.Invoice, error)
type MapEmailClients func() (map[int64]string, error)

// actions
type InsertContainersFn func(containers []entities.Container) error
type InsertOrdersFn func(orders []entities.Order) error
type InsertInvoicesFn func(invoices []entities.Invoice) error
type UpsertBookingFn func(booking entities.Booking) (string, error)
type SetBookingFn func(booking Dto.BookingDto, email string)

// Get operations
type GetContainersByBookingIdFn func(bookingId string) ([]Dto.ContainerDto, error)
type GetOrdersByBookingIdFn func(bookingId string) ([]Dto.OrderDto, error)
type GetContainersByOrderIdFn func(orderId string) ([]Dto.ContainerDto, error)
type GetOrdersByContainerIdFn func(containerId string) ([]Dto.OrderDto, error)
type GetContainersByEmailFn func(email string) ([]Dto.ContainerDto, error)
type GetOrdersByEmailFn func(email string) ([]Dto.OrderDto, error)

// Fallback operations
type FallbackContainersByBookingIdFn func(bookingId string) ([]Dto.ContainerDto, error)
type FallbackOrdersByBookingIdFn func(bookingId string) ([]Dto.OrderDto, error)
type FallbackContainersByOrderIdFn func(orderId string) ([]Dto.ContainerDto, error)
type FallbackOrdersByContainerIdFn func(containerId string) ([]Dto.OrderDto, error)
type FallbackContainersByEmailFn func(email string) ([]Dto.ContainerDto, error)
type FallbackOrdersByEmailFn func(email string) ([]Dto.OrderDto, error)

// Set operations for cache
type SetContainersByBookingIdFn func(bookingId string, containers []Dto.ContainerDto)
type SetOrdersByBookingIdFn func(bookingId string, orders []Dto.OrderDto)
type SetContainersByOrderIdFn func(orderId string, containers []Dto.ContainerDto)
type SetOrdersByContainerIdFn func(containerId string, orders []Dto.OrderDto)
type SetContainersByEmailFn func(email string, containers []Dto.ContainerDto)
type SetOrdersByEmailFn func(email string, orders []Dto.OrderDto)
