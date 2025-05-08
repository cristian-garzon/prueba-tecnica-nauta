package queries

const GetContainers = `
	select 
		container_id, 
		booking_id, 
		container_type, 
		description, 
		weight, 
		created_at 
	from containers 
	where created_at >= $1
`
