package actions

import (
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/types"
)

type GetOrdersByBookingId struct {
	getOrdersByBookingId      types.GetOrdersByBookingIdFn
	fallbackOrdersByBookingId types.FallbackOrdersByBookingIdFn
	setOrdersByBookingId      types.SetOrdersByBookingIdFn
}

func NewGetOrdersByBookingId(
	getOrdersByBookingId types.GetOrdersByBookingIdFn,
	fallbackOrdersByBookingId types.FallbackOrdersByBookingIdFn,
	setOrdersByBookingId types.SetOrdersByBookingIdFn,
) *GetOrdersByBookingId {
	return &GetOrdersByBookingId{
		getOrdersByBookingId:      getOrdersByBookingId,
		fallbackOrdersByBookingId: fallbackOrdersByBookingId,
		setOrdersByBookingId:      setOrdersByBookingId,
	}
}

func (a *GetOrdersByBookingId) Execute(bookingId string) ([]Dto.OrderDto, error) {
	orders, err := a.getOrdersByBookingId(bookingId)
	if err != nil {
		if appErr, ok := err.(*model.AppError); ok && appErr.ErrorCode == model.ErrNoOrdersFound {
			orders, err = a.fallbackOrdersByBookingId(bookingId)
			if err != nil {
				return nil, err
			}
			a.setOrdersByBookingId(bookingId, orders)
		} else {
			return nil, err
		}
	}

	return orders, nil
}
