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
	response := types.NewPaginatedResponse(questions, pagination, len(questions), int(totalPages))

	json.NewEncoder(w).Encode(response)
}

func (server *Server) handleAddQuestion(w http.ResponseWriter, r *http.Request) {
	question := types.NewQuestion()

	err := json.NewDecoder(r.Body).Decode(question)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := types.NewErrorResponse(types.ErrorCodeBadRequest, "")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	isValid, validations := question.Validate()

	if !isValid {
		errorResponse := types.ErrorResponse{}
		for _, validation := range validations {
			errorResponse.Add(types.ErrorCodeBadRequest, validation)
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = server.store.AddQuestion(*question)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := types.NewErrorResponse(types.ErrorCodeDefault, "")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(question)
}
