package api

import (
	"encoding/json"
	"net/http"

	"github.com/erenhncr/go-api-structure/types"
)

func (s *Server) questionsMiddleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		s.handleGetQuestions(w, r)
	case "POST":
		s.handleAddQuestion(w, r)
	default:
		errorResponse := types.NewErrorResponse(types.ErrorCodeUnsupportedMethod, "")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorResponse)
	}
}
