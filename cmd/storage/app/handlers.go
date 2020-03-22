package app

import (
	"github.com/AlisherFozilov/useful/pkg/useful"
	"log"
	"net/http"
)

type ErrorDTO struct {
	err string
}

func (s *Server) handleSaveData() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//request.Body
		err := s.dbSvc.SaveMessagesData()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			err := useful.WriteJSONBody(writer, ErrorDTO{"err.fail"}) // TODO: error
			log.Print(err)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}
