package dbcore

const (
	insertMessageIntoMessages = `INSERT INTO messages (id, type, sender_id, receiver_id, data, timestamp)
VALUES ($1, $2, $3, $4, $5, $6);`
)
