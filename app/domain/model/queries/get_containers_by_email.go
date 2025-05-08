package queries

const GetContainersByEmail = `
	select 
		c.container_id, 
		c.booking_id, 
		c.container_type, 
		c.description, 
		c.weight, 
		c.created_at 
	from containers c
	inner join bookings b on b.booking_id = c.booking_id
	inner join clients cl on cl.id = b.client_id
	where cl.email = $1
`
