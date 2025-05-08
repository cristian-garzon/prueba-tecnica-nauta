package repositories

import (
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/types"
)

type NautaCacheRepository struct {
	bookingCache                 map[string]Dto.BookingDto
	bookingIdsByClientEmailCache map[string][]string
	orderIdByBookingIdCache      map[string]string
	containerIdByBookingIdCache  map[string]string
}

func NewNautaCacheRepository(
	mapBookings types.MapBookings,
	mapContainers types.MapContainers,
	mapOrders types.MapOrders,
	mapInvoices types.MapInvoices,
	mapEmailClients types.MapEmailClients,
) (*NautaCacheRepository, error) {
	bookings, err := mapBookings()
	if err != nil {
		return nil, err
	}

	containers, err := mapContainers()
	if err != nil {
		return nil, err
	}

	orders, err := mapOrders()
	if err != nil {
		return nil, err
	}

	invoices, err := mapInvoices()
	if err != nil {
		return nil, err
	}

	clients, err := mapEmailClients()
	if err != nil {
		return nil, err
	}

	bookingCache := make(map[string]Dto.BookingDto)
	bookingIdsByClientEmailCache := make(map[string][]string)
	orderIdByBookingIdCache := make(map[string]string)
	containerIdByBookingIdCache := make(map[string]string)

	for _, booking := range bookings {
		bookingCache[booking.BookingId] = Dto.FromBooking(
			booking,
			containers[booking.BookingId],
			orders[booking.BookingId],
			invoices,
		)

		for _, order := range orders[booking.BookingId] {
			orderIdByBookingIdCache[order.PurchaseId] = order.BookingId
		}

		for _, container := range containers[booking.BookingId] {
			containerIdByBookingIdCache[container.ContainerId] = container.BookingId
		}

		email := clients[booking.ClientId]
		bookingIdsByClientEmailCache[email] = append(bookingIdsByClientEmailCache[email], booking.BookingId)
	}

	return &NautaCacheRepository{
		bookingCache:                 bookingCache,
		bookingIdsByClientEmailCache: bookingIdsByClientEmailCache,
		orderIdByBookingIdCache:      orderIdByBookingIdCache,
		containerIdByBookingIdCache:  containerIdByBookingIdCache,
	}, nil
}

func (r *NautaCacheRepository) GetContainersByEmail(email string) ([]Dto.ContainerDto, error) {
	bookingIds := r.bookingIdsByClientEmailCache[email]
	if len(bookingIds) == 0 {
		return nil, model.NewAppError("no bookings found", model.ErrNoBookingFound)
	}
	containers := make([]Dto.ContainerDto, 0)
	for _, bookingId := range bookingIds {
		containers = append(containers, r.bookingCache[bookingId].Containers...)
	}
	return containers, nil
}

func (r *NautaCacheRepository) GetOrdersByEmail(email string) ([]Dto.OrderDto, error) {
	bookingIds := r.bookingIdsByClientEmailCache[email]
	if len(bookingIds) == 0 {
		return nil, model.NewAppError("no bookings found", model.ErrNoBookingsFound)
	}
	orders := make([]Dto.OrderDto, 0)
	for _, bookingId := range bookingIds {
		orders = append(orders, r.bookingCache[bookingId].Orders...)
	}
	return orders, nil
}

func (r *NautaCacheRepository) GetOrdersByContainerId(containerId string) ([]Dto.OrderDto, error) {
	bookingId := r.containerIdByBookingIdCache[containerId]
	if bookingId == "" {
		return nil, model.NewAppError("no booking found", model.ErrNoBookingFound)
	}
	return r.bookingCache[bookingId].Orders, nil
}

func (r *NautaCacheRepository) GetOrdersByBookingId(bookingId string) ([]Dto.OrderDto, error) {
	orders := r.bookingCache[bookingId].Orders
	if len(orders) == 0 {
		return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
	}
	return orders, nil
}

func (r *NautaCacheRepository) GetContainersByBookingId(bookingId string) ([]Dto.ContainerDto, error) {
	containers := r.bookingCache[bookingId].Containers
	if len(containers) == 0 {
		return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
	}
	return containers, nil
}

func (r *NautaCacheRepository) GetContainersByOrderId(orderId string) ([]Dto.ContainerDto, error) {
	bookingId := r.orderIdByBookingIdCache[orderId]
	if bookingId == "" {
		return nil, model.NewAppError("no booking found", model.ErrNoBookingFound)
	}
	return r.bookingCache[bookingId].Containers, nil
}

func (r *NautaCacheRepository) SetBooking(booking Dto.BookingDto, email string) {
	r.bookingCache[booking.BookingId] = booking

	for _, order := range booking.Orders {
		r.orderIdByBookingIdCache[order.PurchaseId] = booking.BookingId
	}

	for _, container := range booking.Containers {
		r.containerIdByBookingIdCache[container.ContainerId] = booking.BookingId
	}

	r.bookingIdsByClientEmailCache[email] = append(r.bookingIdsByClientEmailCache[email], booking.BookingId)
}

func (r *NautaCacheRepository) SetContainersByEmail(email string, containers []Dto.ContainerDto) {
	bookingIds := r.bookingIdsByClientEmailCache[email]
	if len(bookingIds) == 0 {
		return
	}

	containersByBooking := make(map[string][]Dto.ContainerDto)
	for _, container := range containers {
		bookingId := r.containerIdByBookingIdCache[container.BookingId]
		if bookingId != "" {
			containersByBooking[bookingId] = append(containersByBooking[bookingId], container)
		}
	}

	for bookingId, containers := range containersByBooking {
		if booking, exists := r.bookingCache[bookingId]; exists {
			booking.Containers = containers
			r.bookingCache[bookingId] = booking
		}
	}
}

func (r *NautaCacheRepository) SetOrdersByEmail(email string, orders []Dto.OrderDto) {
	bookingIds := r.bookingIdsByClientEmailCache[email]
	if len(bookingIds) == 0 {
		return
	}

	ordersByBooking := make(map[string][]Dto.OrderDto)
	for _, order := range orders {
		bookingId := r.orderIdByBookingIdCache[order.BookingId]
		if bookingId != "" {
			ordersByBooking[bookingId] = append(ordersByBooking[bookingId], order)
		}
	}

	for bookingId, orders := range ordersByBooking {
		if booking, exists := r.bookingCache[bookingId]; exists {
			booking.Orders = orders
			r.bookingCache[bookingId] = booking
		}
	}
}

func (r *NautaCacheRepository) SetOrdersByContainerId(containerId string, orders []Dto.OrderDto) {
	bookingId := r.containerIdByBookingIdCache[containerId]
	if bookingId == "" {
		return
	}

	if booking, exists := r.bookingCache[bookingId]; exists {
		booking.Orders = orders
		r.bookingCache[bookingId] = booking
	}
}

func (r *NautaCacheRepository) SetOrdersByBookingId(bookingId string, orders []Dto.OrderDto) {
	if booking, exists := r.bookingCache[bookingId]; exists {
		booking.Orders = orders
		r.bookingCache[bookingId] = booking
	}
}

func (r *NautaCacheRepository) SetContainersByBookingId(bookingId string, containers []Dto.ContainerDto) {
	if booking, exists := r.bookingCache[bookingId]; exists {
		booking.Containers = containers
		r.bookingCache[bookingId] = booking
	}
}

func (r *NautaCacheRepository) SetContainersByOrderId(orderId string, containers []Dto.ContainerDto) {
	bookingId := r.orderIdByBookingIdCache[orderId]
	if bookingId == "" {
		return
	}

	if booking, exists := r.bookingCache[bookingId]; exists {
		booking.Containers = containers
		r.bookingCache[bookingId] = booking
	}
}
