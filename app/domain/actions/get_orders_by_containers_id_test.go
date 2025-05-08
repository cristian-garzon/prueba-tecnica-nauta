package actions

import (
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrdersByContainerId_ok(t *testing.T) {
	action := NewMockGetOrdersByContainerIdOk()

	orders, err := action.Execute("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, orders, []Dto.OrderDto{{PurchaseId: "1"}})
}

func TestGetOrdersByContainerId_fallback_ok(t *testing.T) {
	action := NewMockGetOrdersByContainerIdFallbackOk()

	orders, err := action.Execute("1")

	assert.Nil(t, err)
	assert.Equal(t, orders, []Dto.OrderDto{{PurchaseId: "1"}})
}

func TestGetOrdersByContainerId_fallback_not_found(t *testing.T) {
	action := NewMockGetOrdersByContainerIdFallbackNotFound()

	orders, err := action.Execute("1")

	assert.NotNil(t, err)
	assert.Equal(t, len(orders), 0)
}

func TestGetOrdersByContainerId_UnknownError(t *testing.T) {
	action := NewMockGetOrdersByContainerIdUnknownError()

	orders, err := action.Execute("1")

	assert.NotNil(t, err)
	assert.Equal(t, len(orders), 0)
}
