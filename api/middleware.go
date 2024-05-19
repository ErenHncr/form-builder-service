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
		server.handleGetQuestions(w, r)
	case http.MethodPost:
		server.handleAddQuestion(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK) // TODO: add proper origin checking
	default:
		errorResponse := types.NewErrorResponse(types.ErrorCodeUnsupportedMethod, "")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorResponse)
	}
}
