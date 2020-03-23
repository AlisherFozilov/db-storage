package rpcall

import (
	"context"
	user "github.com/AlisherFozilov/db-storage/pkg/api"
	"github.com/AlisherFozilov/db-storage/pkg/services/dbcore"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Service struct {
	storage *dbcore.Service
	host    string
	port    string
}

func NewService(storage *dbcore.Service, host Host, port Port) *Service {
	return &Service{storage: storage, host: string(host), port: string(port)}
}

type Host string
type Port string

func (s *Service) SaveMessages(ctx context.Context, req *user.Messages) (*user.Response, error) {
	log.Println(req)
	ctxWithTimeOut, _ := context.WithTimeout(ctx, 1*time.Second)
	err := s.storage.SaveMessagesData(ctxWithTimeOut, req)
	if err != nil {
		return &user.Response{Status: 1}, err
	}
	return &user.Response{Status: 0}, nil
}

func (s *Service) StartServing() {
	server := grpc.NewServer()
	user.RegisterStorageServer(server, s)

	listener, err := net.Listen("tcp", net.JoinHostPort(s.host, s.port))
	if err != nil {
		log.Println(err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}
