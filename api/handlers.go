package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/erenhncr/go-api-structure/types"
	"github.com/erenhncr/go-api-structure/util"
)

func (server *Server) handleGetQuestions(w http.ResponseWriter, r *http.Request) {
	pagination := util.GetPagination(r)
	sorting := util.GetSorting(r)

	questions, totalItems, err := server.store.GetQuestions(pagination, sorting)
	if err != nil {
		util.InternalServerError(w, err)
		return
	}

	totalPages := util.GetTotalPages(totalItems, pagination.Size)
	response := types.NewPaginatedResponse(questions, pagination, sorting, totalItems, totalPages)

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
	question := &types.Question{}
	if err := json.NewDecoder(r.Body).Decode(question); err != nil {
		util.BadRequest(w, err)
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

	createdQuestion, err := server.store.CreateQuestion(*question)
	if err != nil {
		util.InternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdQuestion)
}

func (server *Server) handleUpdateQuestion(w http.ResponseWriter, r *http.Request) {
	questionId := r.PathValue("id")

	if questionId == "" {
		errorResponse := types.NewErrorResponse(types.ErrorCodeBadRequest, "question_id_required")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		util.BadRequest(w, err)
		return
	}

	question := &types.Question{}
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

	var questionPatch types.QuestionPatch
	if err = json.Unmarshal(body, &questionPatch); err != nil {
		util.InternalServerError(w, err)
		return
	}

	updatedQuestion, err := server.store.UpdateQuestion(questionId, questionPatch)
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
		util.NotFound(w, err)
		return
	}

	json.NewEncoder(w).Encode(types.ResponseID{
		ID: questionId,
	})
}
