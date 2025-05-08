package Dto

import "prueba-tecnica-nauta/app/domain/model/entities"

type ContainerDto struct {
	ContainerType string  `json:"container_type"`
	ContainerId   string  `json:"container_id"`
	BookingId     string  `json:"booking_id"`
	Description   string  `json:"description"`
	Weight        float64 `json:"weight"`
}

func FromContainer(container entities.Container) ContainerDto {
	return ContainerDto{
		ContainerType: container.ContainerType,
		ContainerId:   container.ContainerId,
		BookingId:     container.BookingId,
		Description:   container.Description,
		Weight:        container.Weight,
	}
}

func (c *ContainerDto) ToContainer(bookingId string) entities.Container {
	return entities.Container{
		ContainerType: c.ContainerType,
		ContainerId:   c.ContainerId,
		Description:   c.Description,
		Weight:        c.Weight,
		BookingId:     bookingId,
	}
}
