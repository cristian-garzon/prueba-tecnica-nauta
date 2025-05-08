package queries

const GetEmailClientId = `
    select email, client_id from clients where created_at >= $1
`
