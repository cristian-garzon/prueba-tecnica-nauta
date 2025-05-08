package actions

import (
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/types"
)

type GetOrdersByEmail struct {
	getOrdersByEmail      types.GetOrdersByEmailFn
	fallbackOrdersByEmail types.FallbackOrdersByEmailFn
	setOrdersByEmail      types.SetOrdersByEmailFn
}

func NewGetOrdersByEmail(
	getOrdersByEmail types.GetOrdersByEmailFn,
	fallbackOrdersByEmail types.FallbackOrdersByEmailFn,
	setOrdersByEmail types.SetOrdersByEmailFn,
) *GetOrdersByEmail {
	return &GetOrdersByEmail{
		getOrdersByEmail:      getOrdersByEmail,
		fallbackOrdersByEmail: fallbackOrdersByEmail,
		setOrdersByEmail:      setOrdersByEmail,
	}
}

func (a *GetOrdersByEmail) Execute(email string) ([]Dto.OrderDto, error) {
	orders, err := a.getOrdersByEmail(email)
	if err != nil {
		if appErr, ok := err.(*model.AppError); ok && appErr.ErrorCode == model.ErrNoOrdersFound {
			orders, err = a.fallbackOrdersByEmail(email)
			if err != nil {
				return nil, err
			}
			a.setOrdersByEmail(email, orders)
		} else {
			return nil, err
		}
	}

	return orders, nil
}
