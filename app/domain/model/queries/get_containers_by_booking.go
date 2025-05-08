package queries

const GetContainersByBooking = `
	select 
		container_id, 
		booking_id, 
		container_type, 
		description, 
		weight, 
		created_at 
	from containers 
	where booking_id = $1
`
