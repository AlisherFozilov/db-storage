package app

import (
	"encoding/json"
	"github.com/AlisherFozilov/useful/pkg/useful"
	"log"
	"net/http"
)

type ErrorDTO struct {
	err string
}

func (s *Server) handleGetAllHistory() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		senderID := request.Context().Value("senderID").(string)
		receiverID := request.Context().Value("receiverID").(string)

		messages, err := s.dbSvc.GetAllMessagesBySenderAndReceiverID(request.Context(), senderID, receiverID)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
			err := useful.WriteJSONBody(writer, ErrorDTO{"err.fail"}) // TODO: error
			if err != nil {
				log.Print(err)
			}
			return
		}

		messagesJSON, err := json.Marshal(messages)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
			err := useful.WriteJSONBody(writer, ErrorDTO{"err.fail"}) // TODO: error
			if err != nil {
				log.Print(err)
			}
			return
		}

		_, err = writer.Write(messagesJSON)
		if err != nil {
			log.Print(err)
			return
		}

		return
	}
}

func (s *Server) handleSaveData() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//request.Body
		err := s.dbSvc.SaveMessagesData(nil, nil)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			err := useful.WriteJSONBody(writer, ErrorDTO{"err.fail"}) // TODO: error
			log.Print(err)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}
