package api

import (
	"encoding/json"
	"net/http"

	"github.com/erenhncr/go-api-structure/storage"
	"github.com/erenhncr/go-api-structure/util"
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
	http.HandleFunc("/user", server.handleGetUserById)
	http.HandleFunc("/user/id", server.handleDeleteUserById)

	http.HandleFunc("/questions", server.questionsMiddleware)

	return http.ListenAndServe(server.listenAddr, nil)
}

func (server *Server) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	user := server.store.Get(10)

	json.NewEncoder(w).Encode(user)
}

func (server *Server) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	user := server.store.Get(10)

	_ = util.Round2Dec(10.34434)

	json.NewEncoder(w).Encode(user)
}

func (server *Server) setDefaultHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*") // TODO: set allowed origins
}
