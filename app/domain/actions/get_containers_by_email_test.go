package actions

import (
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContainersByEmail_ok(t *testing.T) {
	action := NewMockGetContainersByEmailOk()

	containers, err := action.Execute("test@example.com")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, containers, []Dto.ContainerDto{{ContainerId: "1"}})
}

func TestGetContainersByEmail_fallback_ok(t *testing.T) {
	action := NewMockGetContainersByEmailFallbackOk()

	containers, err := action.Execute("test@example.com")

	assert.Nil(t, err)
	assert.Equal(t, containers, []Dto.ContainerDto{{ContainerId: "1"}})
}

func TestGetContainersByEmail_fallback_not_found(t *testing.T) {
	action := NewMockGetContainersByEmailFallbackNotFound()

	containers, err := action.Execute("test@example.com")

	assert.NotNil(t, err)
	assert.Equal(t, len(containers), 0)
}

func TestGetContainersByEmail_UnknownError(t *testing.T) {
	action := NewMockGetContainersByEmailUnknownError()

	containers, err := action.Execute("test@example.com")

	assert.NotNil(t, err)
	assert.Equal(t, len(containers), 0)
}
