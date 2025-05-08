package entities

import (
	"time"
)

type Invoice struct {
	InvoiceId   string
	PurchaseId  string
	Amount      int64
	Status      string
	PaymentDate time.Time
}

func ToMapInvoices(items []map[string]any) map[string][]Invoice {
	invoices := make(map[string][]Invoice)
	for _, item := range items {
		invoice := ToInvoice(item)
		invoices[invoice.PurchaseId] = append(invoices[invoice.PurchaseId], invoice)
	}
	return invoices
}

func ToInvoice(item map[string]any) Invoice {
	return Invoice{
		InvoiceId:   item["invoice_id"].(string),
		PurchaseId:  item["purchase_id"].(string),
		Amount:      item["amount"].(int64),
		Status:      item["status"].(string),
		PaymentDate: item["payment_date"].(time.Time),
	}
}
