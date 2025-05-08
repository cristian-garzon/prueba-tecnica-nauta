package queries

const GetOrdersByContainer = `
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
	inner join containers c on c.booking_id = o.booking_id
	left join invoices i on i.purchase_id = o.purchase_id
	where c.container_id = $1
`
