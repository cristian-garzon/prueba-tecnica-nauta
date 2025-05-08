package actions

import (
	"prueba-tecnica-nauta/app/domain/model"
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/types"
)

type GetContainersByEmail struct {
	getContainersByEmail      types.GetContainersByEmailFn
	fallbackContainersByEmail types.FallbackContainersByEmailFn
	setContainersByEmail      types.SetContainersByEmailFn
}

func NewGetContainersByEmail(
	getContainersByEmail types.GetContainersByEmailFn,
	fallbackContainersByEmail types.FallbackContainersByEmailFn,
	setContainersByEmail types.SetContainersByEmailFn,
) *GetContainersByEmail {
	return &GetContainersByEmail{
		getContainersByEmail:      getContainersByEmail,
		fallbackContainersByEmail: fallbackContainersByEmail,
		setContainersByEmail:      setContainersByEmail,
	}
}

func (a *GetContainersByEmail) Execute(email string) ([]Dto.ContainerDto, error) {
	containers, err := a.getContainersByEmail(email)
	if err != nil {
		if appErr, ok := err.(*model.AppError); ok && appErr.ErrorCode == model.ErrNoContainersFound {
			containers, err = a.fallbackContainersByEmail(email)
			if err != nil {
				return nil, err
			}
			a.setContainersByEmail(email, containers)
		} else {
			return nil, err
		}
	}

	return containers, nil
}
