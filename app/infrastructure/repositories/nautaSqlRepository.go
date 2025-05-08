package repositories

import (
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/entities"
	"prueba-tecnica-nauta/app/domain/model/queries"
	"prueba-tecnica-nauta/app/domain/model/types"
	"prueba-tecnica-nauta/app/infrastructure/circuitbreaker"
	"time"
)

type PostgresRepository struct {
	findAllQuery   types.FindAllQuery
	findOneQuery   types.FindOneQuery
	executeQuery   types.ExecuteQuery
	cacheDuration  time.Time
	circuitBreaker *circuitbreaker.CircuitBreaker
}

func NewPostgresRepository(
	findAllQuery types.FindAllQuery,
	findOneQuery types.FindOneQuery,
	executeQuery types.ExecuteQuery,
	cacheDuration time.Time,
	maxFailures int,
	resetTimeout time.Duration,
) *PostgresRepository {
	return &PostgresRepository{
		findAllQuery:   findAllQuery,
		findOneQuery:   findOneQuery,
		executeQuery:   executeQuery,
		cacheDuration:  cacheDuration,
		circuitBreaker: circuitbreaker.NewCircuitBreaker(maxFailures, resetTimeout),
	}
}

func (r *PostgresRepository) UpsertBooking(bookingDto entities.Booking) (string, error) {
	var response map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		response, err = r.executeQuery(queries.UpsertBooking,
			bookingDto.BookingId,
			bookingDto.ClientId,
			bookingDto.Status,
			bookingDto.OriginPort,
			bookingDto.DestinationPort)
		return err
	}, circuitbreaker.Write)

	if err != nil {
		if err == circuitbreaker.ErrCircuitOpen {
			return "", model.NewAppError("circuit breaker is open", model.ErrQueryError)
		}
		return "", model.NewAppError("error upserting booking"+err.Error(), model.ErrQueryError)
	}
	return response["email"].(string), nil
}

func (r *PostgresRepository) InsertContainers(containers []entities.Container) error {
	containerIds := make([]string, 0, len(containers))
	bookingIds := make([]string, 0, len(containers))
	containerTypes := make([]string, 0, len(containers))
	descriptions := make([]string, 0, len(containers))
	weights := make([]float64, 0, len(containers))

	for _, container := range containers {
		containerIds = append(containerIds, container.ContainerId)
		bookingIds = append(bookingIds, container.BookingId)
		containerTypes = append(containerTypes, container.ContainerType)
		descriptions = append(descriptions, container.Description)
		weights = append(weights, container.Weight)
	}

	var err error
	err = r.circuitBreaker.Execute(func() error {
		_, err = r.executeQuery(queries.InsertContainerBatch, containerIds, bookingIds, containerTypes, descriptions, weights)
		return err
	}, circuitbreaker.Write)

	if err != nil {
		if err == circuitbreaker.ErrCircuitOpen {
			return model.NewAppError("circuit breaker is open", model.ErrQueryError)
		}
		return model.NewAppError("error inserting containers"+err.Error(), model.ErrQueryError)
	}
	return nil
}

func (r *PostgresRepository) InsertOrders(orders []entities.Order) error {
	purchaseIds := make([]string, 0, len(orders))
	bookingIds := make([]string, 0, len(orders))
	statuses := make([]string, 0, len(orders))
	totalAmounts := make([]int64, 0, len(orders))
	descriptions := make([]string, 0, len(orders))

	for _, order := range orders {
		purchaseIds = append(purchaseIds, order.PurchaseId)
		bookingIds = append(bookingIds, order.BookingId)
		statuses = append(statuses, order.Status)
		totalAmounts = append(totalAmounts, order.TotalAmount)
		descriptions = append(descriptions, order.Description)
	}

	var err error
	err = r.circuitBreaker.Execute(func() error {
		_, err = r.executeQuery(queries.InsertOrderBatch, purchaseIds, bookingIds, statuses, totalAmounts, descriptions)
		return err
	}, circuitbreaker.Write)

	if err != nil {
		if err == circuitbreaker.ErrCircuitOpen {
			return model.NewAppError("circuit breaker is open", model.ErrQueryError)
		}
		return model.NewAppError("error inserting orders"+err.Error(), model.ErrQueryError)
	}
	return nil
}

func (r *PostgresRepository) InsertInvoices(invoices []entities.Invoice) error {
	invoiceIds := make([]string, 0, len(invoices))
	purchaseIds := make([]string, 0, len(invoices))
	amounts := make([]int64, 0, len(invoices))
	statuses := make([]string, 0, len(invoices))
	paymentDates := make([]time.Time, 0, len(invoices))

	for _, invoice := range invoices {
		invoiceIds = append(invoiceIds, invoice.InvoiceId)
		purchaseIds = append(purchaseIds, invoice.PurchaseId)
		amounts = append(amounts, invoice.Amount)
		statuses = append(statuses, invoice.Status)
		paymentDates = append(paymentDates, invoice.PaymentDate)
	}

	var err error
	err = r.circuitBreaker.Execute(func() error {
		_, err = r.executeQuery(queries.InsertInvoiceBatch, invoiceIds, purchaseIds, amounts, statuses, paymentDates)
		return err
	}, circuitbreaker.Write)

	if err != nil {
		if err == circuitbreaker.ErrCircuitOpen {
			return model.NewAppError("circuit breaker is open", model.ErrQueryError)
		}
		return model.NewAppError("error inserting invoices"+err.Error(), model.ErrQueryError)
	}
	return nil
}

func (r *PostgresRepository) GetBookings() (map[string]entities.Booking, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetBookings, r.cacheDuration)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no bookings found", model.ErrNoBookingsFound)
		}
		return nil, model.NewAppError("error getting bookings"+err.Error(), model.ErrQueryError)
	}

	return entities.ToMapBookings(items), nil
}

func (r *PostgresRepository) GetContainers() (map[string][]entities.Container, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetContainers, r.cacheDuration)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		}
		return nil, model.NewAppError("error getting containers"+err.Error(), model.ErrQueryError)
	}

	return entities.ToMapContainers(items), nil
}

func (r *PostgresRepository) GetOrders() (map[string][]entities.Order, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetOrders, r.cacheDuration)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		}
		return nil, model.NewAppError("error getting orders"+err.Error(), model.ErrQueryError)
	}

	return entities.ToMapOrders(items), nil
}

func (r *PostgresRepository) GetInvoices() (map[string][]entities.Invoice, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetInvoices, r.cacheDuration)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no invoices found", model.ErrNoOrdersFound)
		}
		return nil, model.NewAppError("error getting invoices"+err.Error(), model.ErrQueryError)
	}

	return entities.ToMapInvoices(items), nil
}

func (r *PostgresRepository) GetEmailClients() (map[int64]string, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetEmailClients, r.cacheDuration)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no clients found", model.ErrNotClientFound)
		}
		return nil, model.NewAppError("error getting email clients"+err.Error(), model.ErrQueryError)
	}

	mapEmailClients := make(map[int64]string)
	for _, item := range items {
		mapEmailClients[item["id"].(int64)] = item["email"].(string)
	}

	return mapEmailClients, nil
}

func (r *PostgresRepository) GetContainersByEmail(email string) ([]Dto.ContainerDto, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetContainersByEmail, email)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		}
		return nil, model.NewAppError("error getting containers"+err.Error(), model.ErrQueryError)
	}

	containers := make([]Dto.ContainerDto, 0, len(items))
	for _, item := range items {
		container := entities.ToContainer(item)
		containers = append(containers, Dto.FromContainer(container))
	}

	return containers, nil
}

func (r *PostgresRepository) GetOrdersByEmail(email string) ([]Dto.OrderDto, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetOrdersByEmail, email)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		}
		return nil, model.NewAppError("error getting orders"+err.Error(), model.ErrQueryError)
	}

	orders := make([]Dto.OrderDto, 0, len(items))
	invoices := make(map[string][]entities.Invoice)

	for _, item := range items {
		order := entities.ToOrder(item)
		if item["invoice_id"] != nil {
			invoice := entities.ToInvoice(item)
			invoices[order.PurchaseId] = append(invoices[order.PurchaseId], invoice)
		}
		orders = append(orders, Dto.FromOrder(order, invoices[order.PurchaseId]))
	}

	return orders, nil
}

func (r *PostgresRepository) GetOrdersByContainerId(containerId string) ([]Dto.OrderDto, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetOrdersByContainer, containerId)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		}
		return nil, model.NewAppError("error getting orders"+err.Error(), model.ErrQueryError)
	}

	orders := make([]Dto.OrderDto, 0, len(items))
	invoices := make(map[string][]entities.Invoice)

	for _, item := range items {
		order := entities.ToOrder(item)
		if item["invoice_id"] != nil {
			invoice := entities.ToInvoice(item)
			invoices[order.PurchaseId] = append(invoices[order.PurchaseId], invoice)
		}
		orders = append(orders, Dto.FromOrder(order, invoices[order.PurchaseId]))
	}

	return orders, nil
}

func (r *PostgresRepository) GetOrdersByBookingId(bookingId string) ([]Dto.OrderDto, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetOrdersByBooking, bookingId)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		}
		return nil, model.NewAppError("error getting orders"+err.Error(), model.ErrQueryError)
	}

	orders := make([]Dto.OrderDto, 0, len(items))
	invoices := make(map[string][]entities.Invoice)

	for _, item := range items {
		order := entities.ToOrder(item)
		if item["invoice_id"] != nil {
			invoice := entities.ToInvoice(item)
			invoices[order.PurchaseId] = append(invoices[order.PurchaseId], invoice)
		}
		orders = append(orders, Dto.FromOrder(order, invoices[order.PurchaseId]))
	}

	return orders, nil
}

func (r *PostgresRepository) GetContainersByBookingId(bookingId string) ([]Dto.ContainerDto, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetContainersByBooking, bookingId)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		}
		return nil, model.NewAppError("error getting containers"+err.Error(), model.ErrQueryError)
	}

	containers := make([]Dto.ContainerDto, 0, len(items))
	for _, item := range items {
		container := entities.ToContainer(item)
		containers = append(containers, Dto.FromContainer(container))
	}

	return containers, nil
}

func (r *PostgresRepository) GetContainersByOrderId(orderId string) ([]Dto.ContainerDto, error) {
	var items []map[string]any
	var err error

	err = r.circuitBreaker.Execute(func() error {
		items, err = r.findAllQuery(queries.GetContainersByOrder, orderId)
		return err
	}, circuitbreaker.Read)

	if err != nil {
		if err == circuitbreaker.ErrCircuitNotFound {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		}
		return nil, model.NewAppError("error getting containers"+err.Error(), model.ErrQueryError)
	}

	containers := make([]Dto.ContainerDto, 0, len(items))
	for _, item := range items {
		container := entities.ToContainer(item)
		containers = append(containers, Dto.FromContainer(container))
	}

	return containers, nil
}
