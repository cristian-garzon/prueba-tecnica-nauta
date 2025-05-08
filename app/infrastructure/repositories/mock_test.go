package repositories

import (
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNautaCacheRepository_GetContainersByEmail(t *testing.T) {
	repo := setupTestRepository()

	t.Run("success", func(t *testing.T) {
		containers, err := repo.GetContainersByEmail("test@example.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(containers))
		assert.Equal(t, "1", containers[0].ContainerId)
	})

	t.Run("not found", func(t *testing.T) {
		containers, err := repo.GetContainersByEmail("nonexistent@example.com")
		assert.NotNil(t, err)
		assert.Equal(t, model.ErrNoBookingFound, err.(*model.AppError).ErrorCode)
		assert.Equal(t, 0, len(containers))
	})
}

func TestNautaCacheRepository_GetOrdersByEmail(t *testing.T) {
	repo := setupTestRepository()

	t.Run("success", func(t *testing.T) {
		orders, err := repo.GetOrdersByEmail("test@example.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(orders))
		assert.Equal(t, "1", orders[0].PurchaseId)
	})

	t.Run("not found", func(t *testing.T) {
		orders, err := repo.GetOrdersByEmail("nonexistent@example.com")
		assert.NotNil(t, err)
		assert.Equal(t, model.ErrNoBookingsFound, err.(*model.AppError).ErrorCode)
		assert.Equal(t, 0, len(orders))
	})
}

func TestNautaCacheRepository_GetOrdersByContainerId(t *testing.T) {
	repo := setupTestRepository()

	t.Run("success", func(t *testing.T) {
		orders, err := repo.GetOrdersByContainerId("1")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(orders))
		assert.Equal(t, "1", orders[0].PurchaseId)
	})

	t.Run("not found", func(t *testing.T) {
		orders, err := repo.GetOrdersByContainerId("nonexistent")
		assert.NotNil(t, err)
		assert.Equal(t, model.ErrNoBookingFound, err.(*model.AppError).ErrorCode)
		assert.Equal(t, 0, len(orders))
	})
}

func TestNautaCacheRepository_GetOrdersByBookingId(t *testing.T) {
	repo := setupTestRepository()

	t.Run("success", func(t *testing.T) {
		orders, err := repo.GetOrdersByBookingId("1")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(orders))
		assert.Equal(t, "1", orders[0].PurchaseId)
	})

	t.Run("not found", func(t *testing.T) {
		orders, err := repo.GetOrdersByBookingId("nonexistent")
		assert.NotNil(t, err)
		assert.Equal(t, model.ErrNoOrdersFound, err.(*model.AppError).ErrorCode)
		assert.Equal(t, 0, len(orders))
	})
}

func TestNautaCacheRepository_GetContainersByBookingId(t *testing.T) {
	repo := setupTestRepository()

	t.Run("success", func(t *testing.T) {
		containers, err := repo.GetContainersByBookingId("1")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(containers))
		assert.Equal(t, "1", containers[0].ContainerId)
	})

	t.Run("not found", func(t *testing.T) {
		containers, err := repo.GetContainersByBookingId("nonexistent")
		assert.NotNil(t, err)
		assert.Equal(t, model.ErrNoContainersFound, err.(*model.AppError).ErrorCode)
		assert.Equal(t, 0, len(containers))
	})
}

func TestNautaCacheRepository_GetContainersByOrderId(t *testing.T) {
	repo := setupTestRepository()

	t.Run("success", func(t *testing.T) {
		containers, err := repo.GetContainersByOrderId("1")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(containers))
		assert.Equal(t, "1", containers[0].ContainerId)
	})

	t.Run("not found", func(t *testing.T) {
		containers, err := repo.GetContainersByOrderId("nonexistent")
		assert.NotNil(t, err)
		assert.Equal(t, model.ErrNoBookingFound, err.(*model.AppError).ErrorCode)
		assert.Equal(t, 0, len(containers))
	})
}

func TestNautaCacheRepository_SetBooking(t *testing.T) {
	repo := setupTestRepository()

	newBooking := Dto.BookingDto{
		BookingId: "2",
		Containers: []Dto.ContainerDto{
			{ContainerId: "2"},
		},
		Orders: []Dto.OrderDto{
			{PurchaseId: "2"},
		},
	}

	repo.SetBooking(newBooking, "test@example.com")

	containers, err := repo.GetContainersByEmail("test@example.com")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(containers))

	orders, err := repo.GetOrdersByEmail("test@example.com")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(orders))
}

func TestNautaCacheRepository_SetContainersByEmail(t *testing.T) {
	repo := setupTestRepository()

	newContainers := []Dto.ContainerDto{
		{ContainerId: "2", BookingId: "1"},
	}
	repo.SetContainersByEmail("test@example.com", newContainers)

	containers, err := repo.GetContainersByEmail("test@example.com")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(containers))
	assert.Equal(t, "2", containers[0].ContainerId)
}

func TestNautaCacheRepository_SetOrdersByEmail(t *testing.T) {
	repo := setupTestRepository()

	newOrders := []Dto.OrderDto{
		{PurchaseId: "2", BookingId: "1"},
	}
	repo.SetOrdersByEmail("test@example.com", newOrders)

	orders, err := repo.GetOrdersByEmail("test@example.com")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(orders))
	assert.Equal(t, "2", orders[0].PurchaseId)
}

func TestNautaCacheRepository_SetOrdersByContainerId(t *testing.T) {
	repo := setupTestRepository()

	newOrders := []Dto.OrderDto{
		{PurchaseId: "2"},
	}
	repo.SetOrdersByContainerId("1", newOrders)

	orders, err := repo.GetOrdersByContainerId("1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(orders))
	assert.Equal(t, "2", orders[0].PurchaseId)
}

func TestNautaCacheRepository_SetOrdersByBookingId(t *testing.T) {
	repo := setupTestRepository()

	newOrders := []Dto.OrderDto{
		{PurchaseId: "2"},
	}
	repo.SetOrdersByBookingId("1", newOrders)

	orders, err := repo.GetOrdersByBookingId("1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(orders))
	assert.Equal(t, "2", orders[0].PurchaseId)
}

func TestNautaCacheRepository_SetContainersByBookingId(t *testing.T) {
	repo := setupTestRepository()

	newContainers := []Dto.ContainerDto{
		{ContainerId: "2"},
	}
	repo.SetContainersByBookingId("1", newContainers)

	containers, err := repo.GetContainersByBookingId("1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(containers))
	assert.Equal(t, "2", containers[0].ContainerId)
}

func TestNautaCacheRepository_SetContainersByOrderId(t *testing.T) {
	repo := setupTestRepository()

	newContainers := []Dto.ContainerDto{
		{ContainerId: "2"},
	}
	repo.SetContainersByOrderId("1", newContainers)

	containers, err := repo.GetContainersByOrderId("1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(containers))
	assert.Equal(t, "2", containers[0].ContainerId)
}

func setupTestRepository() *NautaCacheRepository {
	booking := entities.Booking{
		BookingId: "1",
		ClientId:  1,
	}

	container := entities.Container{
		ContainerId: "1",
		BookingId:   "1",
	}

	order := entities.Order{
		PurchaseId: "1",
		BookingId:  "1",
	}

	invoice := entities.Invoice{
		InvoiceId:  "1",
		PurchaseId: "1",
	}

	repo, _ := NewNautaCacheRepository(
		func() (map[string]entities.Booking, error) {
			return map[string]entities.Booking{
				"1": booking,
			}, nil
		},
		func() (map[string][]entities.Container, error) {
			return map[string][]entities.Container{
				"1": {container},
			}, nil
		},
		func() (map[string][]entities.Order, error) {
			return map[string][]entities.Order{
				"1": {order},
			}, nil
		},
		func() (map[string][]entities.Invoice, error) {
			return map[string][]entities.Invoice{
				"1": {invoice},
			}, nil
		},
		func() (map[int64]string, error) {
			return map[int64]string{
				1: "test@example.com",
			}, nil
		},
	)

	return repo
}
