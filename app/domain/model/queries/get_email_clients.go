package queries

const GetEmailClients = `
    select email, id from clients where created_at >= $1
`
