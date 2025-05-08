package actions

import (
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
)

func NewMockGetContainersByBookingIdOk() *GetContainersByBookingId {
	return &GetContainersByBookingId{
		getContainersByBookingId: func(bookingId string) ([]Dto.ContainerDto, error) {
			return []Dto.ContainerDto{{ContainerId: "1"}}, nil
		},
		fallbackContainersByBookingId: func(bookingId string) ([]Dto.ContainerDto, error) {
			return nil, nil
		},
		setContainersByBookingId: func(bookingId string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetContainersByBookingIdFallbackOk() *GetContainersByBookingId {
	return &GetContainersByBookingId{
		getContainersByBookingId: func(bookingId string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		},
		fallbackContainersByBookingId: func(bookingId string) ([]Dto.ContainerDto, error) {
			return []Dto.ContainerDto{{ContainerId: "1"}}, nil
		},
		setContainersByBookingId: func(bookingId string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetContainersByBookingIdFallbackNotFound() *GetContainersByBookingId {
	return &GetContainersByBookingId{
		getContainersByBookingId: func(bookingId string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		},
		fallbackContainersByBookingId: func(bookingId string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		},
		setContainersByBookingId: func(bookingId string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetContainersByBookingIdUnknownError() *GetContainersByBookingId {
	return &GetContainersByBookingId{
		getContainersByBookingId: func(bookingId string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("unknown error", model.UnknownError)
		},
		fallbackContainersByBookingId: func(bookingId string) ([]Dto.ContainerDto, error) {
			return nil, nil
		},
		setContainersByBookingId: func(bookingId string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetOrdersByEmailOk() *GetOrdersByEmail {
	return &GetOrdersByEmail{
		getOrdersByEmail: func(email string) ([]Dto.OrderDto, error) {
			return []Dto.OrderDto{{PurchaseId: "1"}}, nil
		},
		fallbackOrdersByEmail: func(email string) ([]Dto.OrderDto, error) {
			return nil, nil
		},
		setOrdersByEmail: func(email string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByEmailFallbackOk() *GetOrdersByEmail {
	return &GetOrdersByEmail{
		getOrdersByEmail: func(email string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		},
		fallbackOrdersByEmail: func(email string) ([]Dto.OrderDto, error) {
			return []Dto.OrderDto{{PurchaseId: "1"}}, nil
		},
		setOrdersByEmail: func(email string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByEmailFallbackNotFound() *GetOrdersByEmail {
	return &GetOrdersByEmail{
		getOrdersByEmail: func(email string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		},
		fallbackOrdersByEmail: func(email string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		},
		setOrdersByEmail: func(email string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByEmailUnknownError() *GetOrdersByEmail {
	return &GetOrdersByEmail{
		getOrdersByEmail: func(email string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("unknown error", model.UnknownError)
		},
		fallbackOrdersByEmail: func(email string) ([]Dto.OrderDto, error) {
			return nil, nil
		},
		setOrdersByEmail: func(email string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByBookingIdOk() *GetOrdersByBookingId {
	return &GetOrdersByBookingId{
		getOrdersByBookingId: func(bookingId string) ([]Dto.OrderDto, error) {
			return []Dto.OrderDto{{PurchaseId: "1"}}, nil
		},
		fallbackOrdersByBookingId: func(bookingId string) ([]Dto.OrderDto, error) {
			return nil, nil
		},
		setOrdersByBookingId: func(bookingId string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByBookingIdFallbackOk() *GetOrdersByBookingId {
	return &GetOrdersByBookingId{
		getOrdersByBookingId: func(bookingId string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		},
		fallbackOrdersByBookingId: func(bookingId string) ([]Dto.OrderDto, error) {
			return []Dto.OrderDto{{PurchaseId: "1"}}, nil
		},
		setOrdersByBookingId: func(bookingId string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByBookingIdFallbackNotFound() *GetOrdersByBookingId {
	return &GetOrdersByBookingId{
		getOrdersByBookingId: func(bookingId string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		},
		fallbackOrdersByBookingId: func(bookingId string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		},
		setOrdersByBookingId: func(bookingId string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByBookingIdUnknownError() *GetOrdersByBookingId {
	return &GetOrdersByBookingId{
		getOrdersByBookingId: func(bookingId string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("unknown error", model.UnknownError)
		},
		fallbackOrdersByBookingId: func(bookingId string) ([]Dto.OrderDto, error) {
			return nil, nil
		},
		setOrdersByBookingId: func(bookingId string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByContainerIdOk() *GetOrdersByContainerId {
	return &GetOrdersByContainerId{
		getOrdersByContainerId: func(containerId string) ([]Dto.OrderDto, error) {
			return []Dto.OrderDto{{PurchaseId: "1"}}, nil
		},
		fallbackOrdersByContainerId: func(containerId string) ([]Dto.OrderDto, error) {
			return nil, nil
		},
		setOrdersByContainerId: func(containerId string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByContainerIdFallbackOk() *GetOrdersByContainerId {
	return &GetOrdersByContainerId{
		getOrdersByContainerId: func(containerId string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		},
		fallbackOrdersByContainerId: func(containerId string) ([]Dto.OrderDto, error) {
			return []Dto.OrderDto{{PurchaseId: "1"}}, nil
		},
		setOrdersByContainerId: func(containerId string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByContainerIdFallbackNotFound() *GetOrdersByContainerId {
	return &GetOrdersByContainerId{
		getOrdersByContainerId: func(containerId string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		},
		fallbackOrdersByContainerId: func(containerId string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("no orders found", model.ErrNoOrdersFound)
		},
		setOrdersByContainerId: func(containerId string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetOrdersByContainerIdUnknownError() *GetOrdersByContainerId {
	return &GetOrdersByContainerId{
		getOrdersByContainerId: func(containerId string) ([]Dto.OrderDto, error) {
			return nil, model.NewAppError("unknown error", model.UnknownError)
		},
		fallbackOrdersByContainerId: func(containerId string) ([]Dto.OrderDto, error) {
			return nil, nil
		},
		setOrdersByContainerId: func(containerId string, orders []Dto.OrderDto) {},
	}
}

func NewMockGetContainersByOrderIdOk() *GetContainersByOrderId {
	return &GetContainersByOrderId{
		getContainersByOrderId: func(orderId string) ([]Dto.ContainerDto, error) {
			return []Dto.ContainerDto{{ContainerId: "1"}}, nil
		},
		fallbackContainersByOrderId: func(orderId string) ([]Dto.ContainerDto, error) {
			return nil, nil
		},
		setContainersByOrderId: func(orderId string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetContainersByOrderIdFallbackOk() *GetContainersByOrderId {
	return &GetContainersByOrderId{
		getContainersByOrderId: func(orderId string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		},
		fallbackContainersByOrderId: func(orderId string) ([]Dto.ContainerDto, error) {
			return []Dto.ContainerDto{{ContainerId: "1"}}, nil
		},
		setContainersByOrderId: func(orderId string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetContainersByOrderIdFallbackNotFound() *GetContainersByOrderId {
	return &GetContainersByOrderId{
		getContainersByOrderId: func(orderId string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		},
		fallbackContainersByOrderId: func(orderId string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		},
		setContainersByOrderId: func(orderId string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetContainersByOrderIdUnknownError() *GetContainersByOrderId {
	return &GetContainersByOrderId{
		getContainersByOrderId: func(orderId string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("unknown error", model.UnknownError)
		},
		fallbackContainersByOrderId: func(orderId string) ([]Dto.ContainerDto, error) {
			return nil, nil
		},
		setContainersByOrderId: func(orderId string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetContainersByEmailOk() *GetContainersByEmail {
	return &GetContainersByEmail{
		getContainersByEmail: func(email string) ([]Dto.ContainerDto, error) {
			return []Dto.ContainerDto{{ContainerId: "1"}}, nil
		},
		fallbackContainersByEmail: func(email string) ([]Dto.ContainerDto, error) {
			return nil, nil
		},
		setContainersByEmail: func(email string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetContainersByEmailFallbackOk() *GetContainersByEmail {
	return &GetContainersByEmail{
		getContainersByEmail: func(email string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		},
		fallbackContainersByEmail: func(email string) ([]Dto.ContainerDto, error) {
			return []Dto.ContainerDto{{ContainerId: "1"}}, nil
		},
		setContainersByEmail: func(email string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetContainersByEmailFallbackNotFound() *GetContainersByEmail {
	return &GetContainersByEmail{
		getContainersByEmail: func(email string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		},
		fallbackContainersByEmail: func(email string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("no containers found", model.ErrNoContainersFound)
		},
		setContainersByEmail: func(email string, containers []Dto.ContainerDto) {},
	}
}

func NewMockGetContainersByEmailUnknownError() *GetContainersByEmail {
	return &GetContainersByEmail{
		getContainersByEmail: func(email string) ([]Dto.ContainerDto, error) {
			return nil, model.NewAppError("unknown error", model.UnknownError)
		},
		fallbackContainersByEmail: func(email string) ([]Dto.ContainerDto, error) {
			return nil, nil
		},
		setContainersByEmail: func(email string, containers []Dto.ContainerDto) {},
	}
}
