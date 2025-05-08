package builder

import "prueba-tecnica-nauta/app/domain/actions"

type Actions struct {
	AddBookingAction               *actions.AddBookingAction
	GetContainersByBookingIdAction *actions.GetContainersByBookingId
	GetContainersByOrderIdAction   *actions.GetContainersByOrderId
	GetContainersByEmailAction     *actions.GetContainersByEmail
	GetOrdersByBookingIdAction     *actions.GetOrdersByBookingId
	GetOrdersByContainerIdAction   *actions.GetOrdersByContainerId
	GetOrdersByEmailAction         *actions.GetOrdersByEmail
}

func NewActionsBuilder(
	repositoriesBuilder *RepositoriesBuilder,
) *Actions {
	AddBookingAction := actions.NewAddBookingAction(
		repositoriesBuilder.sqlRepository.InsertContainers,
		repositoriesBuilder.sqlRepository.InsertOrders,
		repositoriesBuilder.sqlRepository.InsertInvoices,
		repositoriesBuilder.sqlRepository.UpsertBooking,
		repositoriesBuilder.cacheRepository.SetBooking,
	)

	getContainersByBookingIdAction := actions.NewGetContainersByBookingId(
		repositoriesBuilder.cacheRepository.GetContainersByBookingId,
		repositoriesBuilder.sqlRepository.GetContainersByBookingId,
		repositoriesBuilder.cacheRepository.SetContainersByBookingId,
	)

	getContainersByOrderIdAction := actions.NewGetContainersByOrderId(
		repositoriesBuilder.cacheRepository.GetContainersByOrderId,
		repositoriesBuilder.sqlRepository.GetContainersByOrderId,
		repositoriesBuilder.cacheRepository.SetContainersByOrderId,
	)

	getContainersByEmailAction := actions.NewGetContainersByEmail(
		repositoriesBuilder.cacheRepository.GetContainersByEmail,
		repositoriesBuilder.sqlRepository.GetContainersByEmail,
		repositoriesBuilder.cacheRepository.SetContainersByEmail,
	)

	getOrdersByBookingIdAction := actions.NewGetOrdersByBookingId(
		repositoriesBuilder.cacheRepository.GetOrdersByBookingId,
		repositoriesBuilder.sqlRepository.GetOrdersByBookingId,
		repositoriesBuilder.cacheRepository.SetOrdersByBookingId,
	)

	getOrdersByContainerIdAction := actions.NewGetOrdersByContainerId(
		repositoriesBuilder.cacheRepository.GetOrdersByContainerId,
		repositoriesBuilder.sqlRepository.GetOrdersByContainerId,
		repositoriesBuilder.cacheRepository.SetOrdersByContainerId,
	)

	getOrdersByEmailAction := actions.NewGetOrdersByEmail(
		repositoriesBuilder.cacheRepository.GetOrdersByEmail,
		repositoriesBuilder.sqlRepository.GetOrdersByEmail,
		repositoriesBuilder.cacheRepository.SetOrdersByEmail,
	)

	return &Actions{
		AddBookingAction:               AddBookingAction,
		GetContainersByBookingIdAction: getContainersByBookingIdAction,
		GetContainersByOrderIdAction:   getContainersByOrderIdAction,
		GetContainersByEmailAction:     getContainersByEmailAction,
		GetOrdersByBookingIdAction:     getOrdersByBookingIdAction,
		GetOrdersByContainerIdAction:   getOrdersByContainerIdAction,
		GetOrdersByEmailAction:         getOrdersByEmailAction,
	}
}
