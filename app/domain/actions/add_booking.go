package actions

import (
	"prueba-tecnica-nauta/app/domain/model/Dto"
	"prueba-tecnica-nauta/app/domain/model/entities"
	"prueba-tecnica-nauta/app/domain/model/types"
)

type AddBookingAction struct {
	insertContainersFn types.InsertContainersFn
	insertOrdersFn     types.InsertOrdersFn
	insertInvoicesFn   types.InsertInvoicesFn
	upsertBookingFn    types.UpsertBookingFn
	setBookingFn       types.SetBookingFn
}

func NewAddBookingAction(
	insertContainersFn types.InsertContainersFn,
	insertOrdersFn types.InsertOrdersFn,
	insertInvoicesFn types.InsertInvoicesFn,
	upsertBookingFn types.UpsertBookingFn,
	setBookingFn types.SetBookingFn,
) *AddBookingAction {
	return &AddBookingAction{
		insertContainersFn: insertContainersFn,
		insertOrdersFn:     insertOrdersFn,
		insertInvoicesFn:   insertInvoicesFn,
		upsertBookingFn:    upsertBookingFn,
		setBookingFn:       setBookingFn,
	}
}

func (a *AddBookingAction) Execute(bookingDto Dto.BookingDto) error {
	booking := bookingDto.ToBooking(bookingDto.ClientId)

	email, err := a.upsertBookingFn(booking)
	if err != nil {
		return err
	}

	containers := make([]entities.Container, 0, len(bookingDto.Containers))
	for _, containerDto := range bookingDto.Containers {
		container := containerDto.ToContainer(booking.BookingId)
		containers = append(containers, container)
	}

	err = a.insertContainersFn(containers)

	if err != nil {
		return err
	}

	orders := make([]entities.Order, 0, len(bookingDto.Orders))
	invoices := make([]entities.Invoice, 0, len(bookingDto.Orders))

	for _, orderDto := range bookingDto.Orders {
		order := orderDto.ToOrder(booking.BookingId)
		orders = append(orders, order)
		for _, invoiceDto := range orderDto.Invoices {
			invoice := invoiceDto.ToInvoice(order.PurchaseId)
			invoices = append(invoices, invoice)
		}
	}

	err = a.insertOrdersFn(orders)

	if err != nil {
		return err
	}

	err = a.insertInvoicesFn(invoices)

	if err != nil {
		return err
	}

	a.setBookingFn(bookingDto, email)

	return nil
}
