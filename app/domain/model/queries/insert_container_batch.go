package queries

const InsertContainerBatch = `
INSERT INTO containers (container_id, booking_id, container_type, description, weight)
VALUES (
	unnest($1::text[]),
	unnest($2::text[]),
	unnest($3::text[]),
	unnest($4::text[]),
	unnest($5::float[])
)
RETURNING container_id
`
