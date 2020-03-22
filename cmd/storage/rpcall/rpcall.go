package rpcall

import (
	"context"
	user "db-storage/pkg/api"
	"db-storage/pkg/services/dbcore"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
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
	fmt.Println(req)
	//s.storage.SaveMessagesData()
	return &user.Response{Status: 1}, nil
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
