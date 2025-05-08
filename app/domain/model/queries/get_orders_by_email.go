package queries

const GetOrdersByEmail = `
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
	inner join bookings b on b.booking_id = o.booking_id
	inner join clients cl on cl.id = b.client_id
	left join invoices i on i.purchase_id = o.purchase_id
	where cl.email = $1
`
