package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/AlisherFozilov/adisher/pkg/di"
	"github.com/AlisherFozilov/db-storage/cmd/storage/app"
	"github.com/AlisherFozilov/db-storage/cmd/storage/rpcall"
	"github.com/AlisherFozilov/db-storage/pkg/services/dbcore"
	"github.com/AlisherFozilov/mymux/pkg/exactmux"
	"github.com/jackc/pgx/v4/pgxpool"
	"net"
	"net/http"
)

// flag - max priority, env - lower priority

var (
	httpHost = flag.String("httpHost", "0.0.0.0", "served httpHost")
	httpPort = flag.String("httpPort", "6666", "served httpPort")
	dsn      = flag.String("dsn", "postgres://user:pass@localhost:5432/postgres", "Postgres DSN")
	tcpHost  = flag.String("tcpHost", "0.0.0.0", "served tcpHost")
	tcpPort  = flag.String("tcpPort", "7777", "served tcpPort")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*httpHost, *httpPort)

	start(addr, *dsn)
}

func start(addr string, dsn string) {

	container := di.NewContainer()
	container.Provide(
		func() rpcall.Host { return rpcall.Host(*tcpHost) },
		func() rpcall.Port { return rpcall.Port(*tcpPort) },
		rpcall.NewService,
		exactmux.NewExactMux,
		func() dbcore.DSN { return dbcore.DSN(dsn) },
		func(dsn dbcore.DSN) *pgxpool.Pool {
			pool, err := pgxpool.Connect(context.Background(), string(dsn))
			if err != nil {
				panic(fmt.Errorf("can't create pool: %w", err))
			}
			return pool
		},
		dbcore.NewService,
		app.NewServer,
	)

	container.Start()

	rpcServer := &rpcall.Service{}
	container.Component(&rpcServer)
	go func() {
		rpcServer.StartServing()
	}()

	server := &app.Server{}
	container.Component(&server)

	panic(http.ListenAndServe(addr, server))
}
