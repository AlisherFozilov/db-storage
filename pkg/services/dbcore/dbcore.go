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
	_, err = s.pool.Exec(ctx, insertMessageIntoMessages,
		message.Id,
		message.Type,
		message.SenderId,
		message.ReceiverId,
		message.Data,
		message.Timestamp,
	)
	return err
}
