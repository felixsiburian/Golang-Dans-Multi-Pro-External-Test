package queries

const (
	QueryGetBuyerByEmail = `
		SELECT 
			*
		FROM
			users u
		WHERE
			u.email = $1 AND
			u.is_active = true
	`
)
