package actions

import (
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContainersByBookingId_ok(t *testing.T) {
	action := NewMockGetContainersByBookingIdOk()

	containers, err := action.Execute("1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, containers, []Dto.ContainerDto{{ContainerId: "1"}})
}

func TestGetContainersByBookingId_fallback_ok(t *testing.T) {
	action := NewMockGetContainersByBookingIdFallbackOk()

	containers, err := action.Execute("1")

	assert.Nil(t, err)
	assert.Equal(t, containers, []Dto.ContainerDto{{ContainerId: "1"}})
}

func TestGetContainersByBookingId_fallback_not_found(t *testing.T) {
	action := NewMockGetContainersByBookingIdFallbackNotFound()

	containers, err := action.Execute("1")

	assert.NotNil(t, err)
	assert.Equal(t, len(containers), 0)
}

func TestGetContainersByBookingId_UnknownError(t *testing.T) {
	action := NewMockGetContainersByBookingIdUnknownError()

	containers, err := action.Execute("1")

	assert.NotNil(t, err)
	assert.Equal(t, len(containers), 0)
}
