package actions

import (
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/types"
)

type GetContainersByOrderId struct {
	getContainersByOrderId      types.GetContainersByOrderIdFn
	fallbackContainersByOrderId types.FallbackContainersByOrderIdFn
	setContainersByOrderId      types.SetContainersByOrderIdFn
}

func NewGetContainersByOrderId(
	getContainersByOrderId types.GetContainersByOrderIdFn,
	fallbackContainersByOrderId types.FallbackContainersByOrderIdFn,
	setContainersByOrderId types.SetContainersByOrderIdFn,
) *GetContainersByOrderId {
	return &GetContainersByOrderId{
		getContainersByOrderId:      getContainersByOrderId,
		fallbackContainersByOrderId: fallbackContainersByOrderId,
		setContainersByOrderId:      setContainersByOrderId,
	}
}

func (a *GetContainersByOrderId) Execute(orderId string) ([]Dto.ContainerDto, error) {
	containers, err := a.getContainersByOrderId(orderId)
	if err != nil {
		if appErr, ok := err.(*model.AppError); ok && appErr.ErrorCode == model.ErrNoContainersFound {
			containers, err = a.fallbackContainersByOrderId(orderId)
			if err != nil {
				return nil, err
			}
			a.setContainersByOrderId(orderId, containers)
		} else {
			return nil, err
		}
	}

	return containers, nil
}
