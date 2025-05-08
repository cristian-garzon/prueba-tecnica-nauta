package queries

const GetOrdersByBooking = `
	select 
		o.purchase_id, 
		o.booking_id, 
		o.status, 
		o.total_amount, 
		o.description, 
		o.created_at,
		i.invoice_id,
		i.amount,
		i.status as invoice_status,
		i.payment_date
	from orders o
	left join invoices i on i.purchase_id = o.purchase_id
	where o.booking_id = $1
`
