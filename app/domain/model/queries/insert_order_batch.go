package queries

const InsertOrderBatch = `
    INSERT INTO orders (purchase_id, booking_id, status, total_amount, description)
    VALUES (
    	unnest($1::text[]),
		unnest($2::text[]),
		unnest($3::text[]),
		unnest($4::bigint[]),
		unnest($5::text[])
	)
	RETURNING purchase_id
`
