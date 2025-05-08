package actions

import (
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/types"
)

type GetOrdersByContainerId struct {
	getOrdersByContainerId      types.GetOrdersByContainerIdFn
	fallbackOrdersByContainerId types.FallbackOrdersByContainerIdFn
	setOrdersByContainerId      types.SetOrdersByContainerIdFn
}

func NewGetOrdersByContainerId(
	getOrdersByContainerId types.GetOrdersByContainerIdFn,
	fallbackOrdersByContainerId types.FallbackOrdersByContainerIdFn,
	setOrdersByContainerId types.SetOrdersByContainerIdFn,
) *GetOrdersByContainerId {
	return &GetOrdersByContainerId{
		getOrdersByContainerId:      getOrdersByContainerId,
		fallbackOrdersByContainerId: fallbackOrdersByContainerId,
		setOrdersByContainerId:      setOrdersByContainerId,
	}
}

func (a *GetOrdersByContainerId) Execute(containerId string) ([]Dto.OrderDto, error) {
	orders, err := a.getOrdersByContainerId(containerId)
	if err != nil {
		if appErr, ok := err.(*model.AppError); ok && appErr.ErrorCode == model.ErrNoOrdersFound {
			orders, err = a.fallbackOrdersByContainerId(containerId)
			if err != nil {
				return nil, err
			}
			a.setOrdersByContainerId(containerId, orders)
		} else {
			return nil, err
		}
	}

	return orders, nil
}
