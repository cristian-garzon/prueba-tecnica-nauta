package Dto

import (
	"prueba-tecnica-nauta/app/domain/model/entities"
	"time"
)

type InvoiceDto struct {
	InvoiceId   string    `json:"invoice_id"`
	Amount      int64     `json:"amount"`
	Status      string    `json:"status"`
	PaymentDate time.Time `json:"payment_date"`
}

func FromInvoice(invoice entities.Invoice) InvoiceDto {
	return InvoiceDto{
		InvoiceId:   invoice.InvoiceId,
		Amount:      invoice.Amount,
		Status:      invoice.Status,
		PaymentDate: invoice.PaymentDate,
	}
}

func (i *InvoiceDto) ToInvoice(purchaseId string) entities.Invoice {
	return entities.Invoice{
		PurchaseId:  purchaseId,
		InvoiceId:   i.InvoiceId,
		Amount:      i.Amount,
		Status:      i.Status,
		PaymentDate: i.PaymentDate,
	}
}
