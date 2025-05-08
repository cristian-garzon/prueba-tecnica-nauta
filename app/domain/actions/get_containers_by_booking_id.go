package actions

import (
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/types"
)

type GetContainersByBookingId struct {
	getContainersByBookingId      types.GetContainersByBookingIdFn
	fallbackContainersByBookingId types.FallbackContainersByBookingIdFn
	setContainersByBookingId      types.SetContainersByBookingIdFn
}

func NewGetContainersByBookingId(
	getContainersByBookingId types.GetContainersByBookingIdFn,
	fallbackContainersByBookingId types.FallbackContainersByBookingIdFn,
	setContainersByBookingId types.SetContainersByBookingIdFn,
) *GetContainersByBookingId {
	return &GetContainersByBookingId{
		getContainersByBookingId:      getContainersByBookingId,
		fallbackContainersByBookingId: fallbackContainersByBookingId,
		setContainersByBookingId:      setContainersByBookingId,
	}
}

func (a *GetContainersByBookingId) Execute(bookingId string) ([]Dto.ContainerDto, error) {
	containers, err := a.getContainersByBookingId(bookingId)
	if err != nil {
		if appErr, ok := err.(*model.AppError); ok && appErr.ErrorCode == model.ErrNoContainersFound {
			containers, err = a.fallbackContainersByBookingId(bookingId)
			if err != nil {
				return nil, err
			}
			a.setContainersByBookingId(bookingId, containers)
		} else {
			return nil, err
		}
	}

	return containers, nil
}
