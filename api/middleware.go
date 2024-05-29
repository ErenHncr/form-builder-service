package api

import (
	"encoding/json"
	"net/http"

	"github.com/erenhncr/go-api-structure/types"
)

func (server *Server) questionsMiddleware(w http.ResponseWriter, r *http.Request) {
	server.setDefaultHeaders(&w)

	switch r.Method {
	case http.MethodGet:
		if id := r.PathValue("id"); id != "" {
			server.handleGetQuestion(w, r)
			return
		}
		server.handleGetQuestions(w, r)
	case http.MethodPost:
		server.handleCreateQuestion(w, r)
	case http.MethodPatch:
		server.handleUpdateQuestion(w, r)
	case http.MethodDelete:
		server.handleDeleteQuestion(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK) // TODO: add proper origin checking
	default:
		errorResponse := types.NewErrorResponse(types.ErrorCodeUnsupportedMethod, "")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorResponse)
	}
}
