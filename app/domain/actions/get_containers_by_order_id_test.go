package actions

import (
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContainersByOrderId_ok(t *testing.T) {
	action := NewMockGetContainersByOrderIdOk()

	containers, err := action.Execute("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, containers, []Dto.ContainerDto{{ContainerId: "1"}})
}

func TestGetContainersByOrderId_fallback_ok(t *testing.T) {
	action := NewMockGetContainersByOrderIdFallbackOk()

	containers, err := action.Execute("1")

	assert.Nil(t, err)
	assert.Equal(t, containers, []Dto.ContainerDto{{ContainerId: "1"}})
}

func TestGetContainersByOrderId_fallback_not_found(t *testing.T) {
	action := NewMockGetContainersByOrderIdFallbackNotFound()

	containers, err := action.Execute("1")

	assert.NotNil(t, err)
	assert.Equal(t, len(containers), 0)
}

func TestGetContainersByOrderId_UnknownError(t *testing.T) {
	action := NewMockGetContainersByOrderIdUnknownError()

	containers, err := action.Execute("1")

	assert.NotNil(t, err)
	assert.Equal(t, len(containers), 0)
}
