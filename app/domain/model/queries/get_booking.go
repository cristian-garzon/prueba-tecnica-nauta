package queries

const GetBookings = `
    select 
		booking_id, 
		client_id, 
		status, 
		origin_port, 
		destination_port, 
		created_at, 
		updated_at 
	from bookings where created_at >= $1
`
