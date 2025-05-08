package queries

const GetContainersByOrder = `
	select 
		c.container_id, 
		c.booking_id, 
		c.container_type, 
		c.description, 
		c.weight, 
		c.created_at 
	from containers c
	inner join orders o on o.booking_id = c.booking_id
	where o.purchase_id = $1
`
