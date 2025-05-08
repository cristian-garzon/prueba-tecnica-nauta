package queries

const GetOrders = `
    select 
		purchase_id, 
		booking_id, 
		status, 
		total_amount, 
		description, 
		created_at 
	from orders 
	where created_at >= $1
`
