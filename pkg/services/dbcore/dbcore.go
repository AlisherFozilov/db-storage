package dbcore

import (
	"context"
	user "github.com/AlisherFozilov/db-storage/pkg/api"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DSN string

type Service struct {
	pool *pgxpool.Pool
	dsn  string
}

func NewService(pool *pgxpool.Pool, dsn DSN) *Service {
	return &Service{pool: pool, dsn: string(dsn)}
}

func (s *Service) Start() {
	ddls := []string{messagesDDL}
	for _, ddl := range ddls {
		_, err := s.pool.Exec(context.Background(), ddl)
		if err != nil {
			panic(err)
		}
	}
}

func (s *Service) SaveMessagesData(ctx context.Context, message *user.Messages) (err error) {
	_, err = s.pool.Exec(ctx, insertMessageIntoMessagesSql,
		message.Id,
		message.Type,
		message.SenderId,
		message.ReceiverId,
		message.Data,
		message.Timestamp,
	)
	return err
}

func (s *Service) GetAllMessagesBySenderAndReceiverID(ctx context.Context, senderID string, receiverID string,
) (messages []Message, err error) {
	rows, err := s.pool.Query(ctx, getMessagesBySenderAndReceiverIdSql,
		senderID,
		receiverID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages = make([]Message, 0)
	message := Message{}
	for rows.Next() {
		err := rows.Scan(
			&message.ID,
			&message.Type,
			&message.SenderID,
			&message.ReceiverID,
			&message.Data,
			&message.TimeStamp,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, err
}

type Message struct {
	ID         string `json:"id"`
	Type       int    `json:"type"`
	SenderID   int64  `json:"sender_id"`
	ReceiverID int64  `json:"receiver_id"`
	Data       string `json:"text"`
	TimeStamp  int64  `json:"time_stamp"`
}
