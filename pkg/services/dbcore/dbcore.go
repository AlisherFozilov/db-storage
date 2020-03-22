package dbcore

import (
	"context"
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

func (s *Service) InitDb() {
	ddls := []string{messagesDDL}
	for _, ddl := range ddls {
		_, err := s.pool.Exec(context.Background(), ddl)
		if err != nil {
			panic(err)
		}
	}
}

func (s *Service) SaveMessagesData() (err error) {
	return nil
}
