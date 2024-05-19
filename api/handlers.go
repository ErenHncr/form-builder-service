package api

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/erenhncr/go-api-structure/types"
)

func (server *Server) handleGetQuestions(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = types.DefaultPageNumber
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		sizeInt = types.DefaultPageSize
	}

	pagination := types.NewPagination(pageInt, sizeInt)

	questions := server.store.GetQuestions(pagination)

	if len(questions) == 0 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := types.NewErrorResponse(types.ErrorCodeNotFound, "")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	totalPages := math.Ceil(float64(len(questions)) / float64(pagination.Size))
	response := types.NewPaginatedResponse(questions, pagination, int(totalPages))

	json.NewEncoder(w).Encode(response)
}

func (server *Server) handleAddQuestion(w http.ResponseWriter, r *http.Request) {
	var question types.Question

	err := json.NewDecoder(r.Body).Decode(&question)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := types.NewErrorResponse(types.ErrorCodeBadRequest, "")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	server.store.AddQuestion(question)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}
