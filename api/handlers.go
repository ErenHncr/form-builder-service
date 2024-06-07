package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/erenhncr/go-api-structure/types"
	"github.com/erenhncr/go-api-structure/util"
)

// GetQuestions godoc
// @Router		/questions [get]
// @Summary	List questions
// @Tags		questions
// @Accept		json
// @Produce	json
// @Param		page	query		int	false	"Page number" default(1)
// @Param		size	query		int	false	"Page size" default(10)
// @Param		sort	query		string	false	"Sort order" default(-createdAt)
// @Success	200		{object}	types.PaginatedResponse[types.Question]
// @Failure	404		{object}	types.ErrorResponse
// @Failure	500		{object}	types.ErrorResponse
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

// GetQuestion godoc
// @Router		/questions/{id} [get]
// @Summary	Get question by id
// @Tags		questions
// @Accept		json
// @Produce	json
// @Param		id		path		string	true	"Question ID"
// @Success	200		{object}	types.Question
// @Failure	400		{object}	types.ErrorResponse
// @Failure	404		{object}	types.ErrorResponse
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

// CreateQuestion godoc
// @Router		/questions [post]
// @Summary	Create question
// @Tags		questions
// @Accept		json
// @Produce	json
// @Param		body	body		types.QuestionPatch	true	"Question"
// @Success	201		{object}	types.Question
// @Failure	400		{object}	types.ErrorResponse
// @Failure	500		{object}	types.ErrorResponse
func (server *Server) handleCreateQuestion(w http.ResponseWriter, r *http.Request) {
	question := &types.Question{}
	if err := json.NewDecoder(r.Body).Decode(question); err != nil {
		util.BadRequest(w, err)
		return
	}

	question.ID = ""
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

// UpdateQuestion godoc
// @Router		/questions/{id} [patch]
// @Summary	Update question by id
// @Tags		questions
// @Accept		json
// @Produce	json
// @Param		id		path		string	true	"Question ID"
// @Param		body	body		types.QuestionPatch	true	"Question Patch"
// @Success	200		{object}	types.Question
// @Failure	400		{object}	types.ErrorResponse
// @Failure	500		{object}	types.ErrorResponse
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

// DeleteQuestion godoc
// @Router		/questions/{id} [delete]
// @Summary	Delete question by id
// @Tags		questions
// @Accept		json
// @Produce	json
// @Param		id		path		string	true	"Question ID"
// @Success	200		{object}	types.ResponseID
// @Failure	400		{object}	types.ErrorResponse
// @Failure	404		{object}	types.ErrorResponse
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
