package actions

import (
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrdersByBookingId_ok(t *testing.T) {
	action := NewMockGetOrdersByBookingIdOk()

	orders, err := action.Execute("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, orders, []Dto.OrderDto{{PurchaseId: "1"}})
}

func TestGetOrdersByBookingId_fallback_ok(t *testing.T) {
	action := NewMockGetOrdersByBookingIdFallbackOk()

	orders, err := action.Execute("1")

	assert.Nil(t, err)
	assert.Equal(t, orders, []Dto.OrderDto{{PurchaseId: "1"}})
}

func TestGetOrdersByBookingId_fallback_not_found(t *testing.T) {
	action := NewMockGetOrdersByBookingIdFallbackNotFound()

	orders, err := action.Execute("1")

	assert.NotNil(t, err)
	assert.Equal(t, len(orders), 0)
}

func TestGetOrdersByBookingId_UnknownError(t *testing.T) {
	action := NewMockGetOrdersByBookingIdUnknownError()

	orders, err := action.Execute("1")

	assert.NotNil(t, err)
	assert.Equal(t, len(orders), 0)
}
