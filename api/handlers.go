package api

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/erenhncr/go-api-structure/types"
	"github.com/erenhncr/go-api-structure/util"
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

func (server *Server) handleGetQuestion(w http.ResponseWriter, r *http.Request) {
	questionId := r.PathValue("id")

	if questionId == "" {
		errorResponse := types.NewErrorResponse(types.ErrorCodeBadRequest, "question_id_required")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	question, err := server.store.GetQuestion(questionId)
	if err != nil {
		errorResponse := types.NewErrorResponse(types.ErrorCodeNotFound, err.Error())
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	json.NewEncoder(w).Encode(question)
}

func (server *Server) handleCreateQuestion(w http.ResponseWriter, r *http.Request) {
	question := types.NewQuestion()

	err := json.NewDecoder(r.Body).Decode(question)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := types.NewErrorResponse(types.ErrorCodeBadRequest, "")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	isValid, validations := question.Validate([]string{})

	if !isValid {
		errorResponse := types.ErrorResponse{}
		for _, validation := range validations {
			errorResponse.Add(types.ErrorCodeBadRequest, validation)
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = server.store.CreateQuestion(*question)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := types.NewErrorResponse(types.ErrorCodeDefault, "")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(question)
}

func (server *Server) handleUpdateQuestion(w http.ResponseWriter, r *http.Request) {
	questionId := r.PathValue("id")

	if questionId == "" {
		errorResponse := types.NewErrorResponse(types.ErrorCodeBadRequest, "question_id_required")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	question := &types.Question{}
	body, _ := io.ReadAll(r.Body)
	bodyKeys := util.GetResponseBodyKeys(body)

	if err := json.Unmarshal(body, question); err != nil {
		errorResponse := types.NewErrorResponse(types.ErrorCodeBadRequest, "invalid_json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if isValid, validations := question.Validate(bodyKeys); !isValid {
		errorResponse := types.ErrorResponse{}
		for _, validation := range validations {
			errorResponse.Add(types.ErrorCodeBadRequest, validation)
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	updatedQuestion, err := server.store.UpdateQuestion(questionId, *question)
	if err != nil {
		errorResponse := types.NewErrorResponse(types.ErrorCodeBadRequest, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	json.NewEncoder(w).Encode(updatedQuestion)
}

func (server *Server) handleDeleteQuestion(w http.ResponseWriter, r *http.Request) {
	questionId := r.PathValue("id")

	if questionId == "" {
		errorResponse := types.NewErrorResponse(types.ErrorCodeBadRequest, "question_id_required")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if err := server.store.DeleteQuestion(questionId); err != nil {
		errorResponse := types.NewErrorResponse(types.ErrorCodeNotFound, err.Error())
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	json.NewEncoder(w).Encode(types.ResponseID{
		ID: questionId,
	})
}
