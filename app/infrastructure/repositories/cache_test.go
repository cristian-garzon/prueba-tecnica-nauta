package repositories

import (
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNautaCacheRepository_GetContainersByEmail_ok(t *testing.T) {
	repo := setupTestRepository()

	containers, err := repo.GetContainersByEmail("test@example.com")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, containers, []Dto.ContainerDto{{ContainerId: "1", BookingId: "1"}})
}

func TestNautaCacheRepository_GetContainersByEmail_not_found(t *testing.T) {
	repo := setupTestRepository()

	containers, err := repo.GetContainersByEmail("nonexistent@example.com")

	assert.NotNil(t, err)
	assert.Equal(t, len(containers), 0)
}

func TestNautaCacheRepository_GetOrdersByEmail_ok(t *testing.T) {
	repo := setupTestRepository()

	orders, err := repo.GetOrdersByEmail("test@example.com")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, orders, []Dto.OrderDto{{PurchaseId: "1", BookingId: "1", Invoices: []Dto.InvoiceDto{{InvoiceId: "1"}}}})
}

func TestNautaCacheRepository_GetOrdersByEmail_not_found(t *testing.T) {
	repo := setupTestRepository()

	orders, err := repo.GetOrdersByEmail("nonexistent@example.com")

	assert.NotNil(t, err)
	assert.Equal(t, len(orders), 0)
}

func TestNautaCacheRepository_GetOrdersByContainerId_ok(t *testing.T) {
	repo := setupTestRepository()

	orders, err := repo.GetOrdersByContainerId("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, orders, []Dto.OrderDto{{PurchaseId: "1", BookingId: "1", Invoices: []Dto.InvoiceDto{{InvoiceId: "1"}}}})
}

func TestNautaCacheRepository_GetOrdersByContainerId_not_found(t *testing.T) {
	repo := setupTestRepository()

	orders, err := repo.GetOrdersByContainerId("nonexistent")

	assert.NotNil(t, err)
	assert.Equal(t, len(orders), 0)
}

func TestNautaCacheRepository_GetOrdersByBookingId_ok(t *testing.T) {
	repo := setupTestRepository()

	orders, err := repo.GetOrdersByBookingId("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, orders, []Dto.OrderDto{{
		PurchaseId: "1",
		BookingId:  "1",
		Invoices:   []Dto.InvoiceDto{{InvoiceId: "1"}},
	}})
}

func TestNautaCacheRepository_GetOrdersByBookingId_not_found(t *testing.T) {
	repo := setupTestRepository()

	orders, err := repo.GetOrdersByBookingId("nonexistent")

	assert.NotNil(t, err)
	assert.Equal(t, len(orders), 0)
}

func TestNautaCacheRepository_GetContainersByBookingId_ok(t *testing.T) {
	repo := setupTestRepository()

	containers, err := repo.GetContainersByBookingId("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, containers, []Dto.ContainerDto{{ContainerId: "1", BookingId: "1"}})
}

func TestNautaCacheRepository_GetContainersByBookingId_not_found(t *testing.T) {
	repo := setupTestRepository()

	containers, err := repo.GetContainersByBookingId("nonexistent")

	assert.NotNil(t, err)
	assert.Equal(t, len(containers), 0)
}

func TestNautaCacheRepository_GetContainersByOrderId_ok(t *testing.T) {
	repo := setupTestRepository()

	containers, err := repo.GetContainersByOrderId("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, containers, []Dto.ContainerDto{{ContainerId: "1", BookingId: "1"}})
}

func TestNautaCacheRepository_GetContainersByOrderId_not_found(t *testing.T) {
	repo := setupTestRepository()

	containers, err := repo.GetContainersByOrderId("nonexistent")

	assert.NotNil(t, err)
	assert.Equal(t, len(containers), 0)
}

func TestNautaCacheRepository_SetBooking_ok(t *testing.T) {
	repo := setupTestRepository()

	newBooking := Dto.BookingDto{
		BookingId: "2",
		Containers: []Dto.ContainerDto{
			{ContainerId: "2", BookingId: "2"},
		},
		Orders: []Dto.OrderDto{
			{PurchaseId: "2", BookingId: "2"},
		},
	}

	repo.SetBooking(newBooking, "test@example.com")

	containers, err := repo.GetContainersByEmail("test@example.com")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, 2, len(containers))
	assert.Equal(t, "2", containers[1].ContainerId)

	orders, err := repo.GetOrdersByEmail("test@example.com")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, 2, len(orders))
	assert.Equal(t, "2", orders[1].PurchaseId)
}

func TestNautaCacheRepository_SetContainersByEmail_ok(t *testing.T) {
	repo := setupTestRepository()

	newContainers := []Dto.ContainerDto{
		{ContainerId: "2", BookingId: "1"},
	}

	repo.SetContainersByEmail("test@example.com", newContainers)

	containers, err := repo.GetContainersByEmail("test@example.com")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, 1, len(containers))
	assert.Equal(t, "2", containers[0].ContainerId)
}

func TestNautaCacheRepository_SetOrdersByEmail_ok(t *testing.T) {
	repo := setupTestRepository()

	newOrders := []Dto.OrderDto{
		{PurchaseId: "2", BookingId: "1"},
	}

	repo.SetOrdersByEmail("test@example.com", newOrders)

	orders, err := repo.GetOrdersByEmail("test@example.com")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, 1, len(orders))
	assert.Equal(t, "2", orders[0].PurchaseId)
}

func TestNautaCacheRepository_SetOrdersByContainerId_ok(t *testing.T) {
	repo := setupTestRepository()

	newOrders := []Dto.OrderDto{
		{PurchaseId: "2", BookingId: "2"},
	}

	repo.SetOrdersByContainerId("1", newOrders)

	orders, err := repo.GetOrdersByContainerId("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, 1, len(orders))
	assert.Equal(t, "2", orders[0].PurchaseId)
}

func TestNautaCacheRepository_SetOrdersByBookingId_ok(t *testing.T) {
	repo := setupTestRepository()

	newOrders := []Dto.OrderDto{
		{PurchaseId: "2", BookingId: "2"},
	}

	repo.SetOrdersByBookingId("1", newOrders)

	orders, err := repo.GetOrdersByBookingId("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, 1, len(orders))
	assert.Equal(t, "2", orders[0].PurchaseId)
}

func TestNautaCacheRepository_SetContainersByBookingId_ok(t *testing.T) {
	repo := setupTestRepository()

	newContainers := []Dto.ContainerDto{
		{ContainerId: "2", BookingId: "2"},
	}

	repo.SetContainersByBookingId("1", newContainers)

	containers, err := repo.GetContainersByBookingId("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, 1, len(containers))
	assert.Equal(t, "2", containers[0].ContainerId)
}

func TestNautaCacheRepository_SetContainersByOrderId_ok(t *testing.T) {
	repo := setupTestRepository()

	newContainers := []Dto.ContainerDto{
		{ContainerId: "2", BookingId: "2"},
	}

	repo.SetContainersByOrderId("1", newContainers)

	containers, err := repo.GetContainersByOrderId("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, 1, len(containers))
	assert.Equal(t, "2", containers[0].ContainerId)
}
