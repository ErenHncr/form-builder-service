package api

import (
	"net/http"

	"github.com/erenhncr/go-api-structure/storage"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (server *Server) Start() error {
	http.HandleFunc("/questions/{id}", server.questionsMiddleware)
	http.HandleFunc("/questions", server.questionsMiddleware)

	return http.ListenAndServe(server.listenAddr, nil)
}

func (server *Server) setDefaultHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*") // TODO: set allowed origins
}
