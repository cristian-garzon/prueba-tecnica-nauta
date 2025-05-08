package entities

import (
	"time"
)

type Container struct {
	ContainerId   string
	BookingId     string
	ContainerType string
	Description   string
	Weight        float64
	CreatedAt     time.Time
}

func ToMapContainers(items []map[string]any) map[string][]Container {
	containers := make(map[string][]Container)
	for _, item := range items {
		container := ToContainer(item)
		containers[container.BookingId] = append(containers[container.BookingId], container)
	}
	return containers
}

func ToContainer(item map[string]any) Container {
	return Container{
		ContainerId:   item["container_id"].(string),
		BookingId:     item["booking_id"].(string),
		ContainerType: item["container_type"].(string),
		Description:   item["description"].(string),
		Weight:        item["weight"].(float64),
		CreatedAt:     item["created_at"].(time.Time),
	}
}
