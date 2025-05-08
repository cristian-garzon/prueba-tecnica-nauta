package queries

const InsertInvoiceBatch = `
INSERT INTO invoices (invoice_id, purchase_id, amount, status, payment_date)
    VALUES (
        unnest($1::text[]),
        unnest($2::text[]),
        unnest($3::bigint[]),
        unnest($4::text[]),
        unnest($5::timestamp[])
    )
	RETURNING invoice_id
`
