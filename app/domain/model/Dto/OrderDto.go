package Dto

import "prueba-tecnica-nauta/app/domain/model/entities"

type OrderDto struct {
	PurchaseId  string       `json:"purchase_id"`
	BookingId   string       `json:"booking_id"`
	Status      string       `json:"status"`
	TotalAmount int64        `json:"total_amount"`
	Description string       `json:"description"`
	Invoices    []InvoiceDto `json:"invoices"`
}

func FromOrder(order entities.Order, invoices []entities.Invoice) OrderDto {
	orderInvoices := make([]InvoiceDto, len(invoices))
	for i, invoice := range invoices {
		if invoice.PurchaseId == order.PurchaseId {
			orderInvoices[i] = FromInvoice(invoice)
		}
	}
	return OrderDto{
		PurchaseId:  order.PurchaseId,
		BookingId:   order.BookingId,
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		Description: order.Description,
		Invoices:    orderInvoices,
	}
}

func (o *OrderDto) ToOrder(bookingId string) entities.Order {
	return entities.Order{
		PurchaseId:  o.PurchaseId,
		Status:      o.Status,
		TotalAmount: o.TotalAmount,
		Description: o.Description,
		BookingId:   bookingId,
	}
}
