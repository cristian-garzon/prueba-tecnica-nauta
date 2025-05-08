package queries

const UpsertBooking = `
INSERT INTO bookings (booking_id, client_id, status, origin_port, destination_port)
    VALUES ($1, $2, $3, $4, $5)
    ON CONFLICT (booking_id) DO UPDATE SET
        status = $3
    RETURNING (SELECT email FROM clients WHERE id = bookings.client_id)
`
