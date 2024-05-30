package util

import (
	"encoding/json"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/erenhncr/go-api-structure/types"
)

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

func GetDatabaseName() string {
	return os.Getenv("DATABASE_NAME")
}

func GetResponseBodyKeys(bodyBytes []byte) []string {
	jsonMap := make(map[string]json.RawMessage)
	err := json.Unmarshal(bodyBytes, &jsonMap)

	keys := make([]string, 0)
	if err != nil {
		return keys
	}

	for key := range jsonMap {
		keys = append(keys, key)
	}

	return keys
}

func GetPagination(r *http.Request) types.Pagination {
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

	return pagination
}

func GetTotalPages(totalItems int, size int) int {
	totalPages := math.Ceil(float64(totalItems) / float64(size))
	return int(totalPages)
}

func InternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	errorResponse := types.NewErrorResponse(types.ErrorCodeInternalServerError, err.Error())
	json.NewEncoder(w).Encode(errorResponse)
}

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	errorResponse := types.NewErrorResponse(types.ErrorCodeBadRequest, err.Error())
	json.NewEncoder(w).Encode(errorResponse)
}

func NotFound(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	errorResponse := types.NewErrorResponse(types.ErrorCodeNotFound, err.Error())
	json.NewEncoder(w).Encode(errorResponse)
}
