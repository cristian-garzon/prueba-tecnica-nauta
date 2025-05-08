package queries

const GetInvoices = `
	select 
		invoice_id, 
		purchase_id, 
		amount, 
		status, 
		payment_date, 
		created_at, 
		updated_at 
	from invoices where created_at >= $1
`
