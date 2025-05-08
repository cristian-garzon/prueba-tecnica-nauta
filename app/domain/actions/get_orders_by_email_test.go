package actions

import (
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrdersByEmail_ok(t *testing.T) {
	action := NewMockGetOrdersByEmailOk()

	orders, err := action.Execute("test@example.com")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, orders, []Dto.OrderDto{{PurchaseId: "1"}})
}

func TestGetOrdersByEmail_fallback_ok(t *testing.T) {
	action := NewMockGetOrdersByEmailFallbackOk()

	orders, err := action.Execute("test@example.com")

	assert.Nil(t, err)
	assert.Equal(t, orders, []Dto.OrderDto{{PurchaseId: "1"}})
}

func TestGetOrdersByEmail_fallback_not_found(t *testing.T) {
	action := NewMockGetOrdersByEmailFallbackNotFound()

	orders, err := action.Execute("test@example.com")

	assert.NotNil(t, err)
	assert.Equal(t, len(orders), 0)
}

func TestGetOrdersByEmail_UnknownError(t *testing.T) {
	action := NewMockGetOrdersByEmailUnknownError()

	orders, err := action.Execute("test@example.com")

	assert.NotNil(t, err)
	assert.Equal(t, len(orders), 0)
}
