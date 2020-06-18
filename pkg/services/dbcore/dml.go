package dbcore

const (
	insertMessageIntoMessagesSql = `INSERT INTO messages (id, type, sender_id, receiver_id, data, timestamp)
VALUES ($1, $2, $3, $4, $5, $6);`

	getMessagesBySenderAndReceiverIdSql = `SELECT id, type, sender_id, receiver_id, data, timestamp FROM messages
WHERE sender_id = $1 OR receiver_id = $2 ORDER BY timestamp ASC;`
)
