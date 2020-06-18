package app

import (
	"github.com/AlisherFozilov/db-storage/pkg/services/dbcore"
	"github.com/AlisherFozilov/mymux/pkg/exactmux"
	"net/http"
)

type Server struct {
	router *exactmux.ExactMux
	dbSvc  *dbcore.Service
}

func NewServer(router *exactmux.ExactMux, dbcoreSvc *dbcore.Service) *Server {
	return &Server{router: router, dbSvc: dbcoreSvc}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Start() {
	s.InitRoutes()
}

func (s *Server) InitRoutes() {
	mux := s.router

	mux.GET("/msg/all_history/{senderID}/{receiverID}", s.handleGetAllHistory())
	//mux.POST("/save/messages", s.handleSaveData()) //bad idea
}
