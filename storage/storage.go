package storage

import "github.com/erenhncr/go-api-structure/types"

type Storage interface {
	Get(int) *types.User
	GetQuestions(types.Pagination) []types.Question
	AddQuestion(types.Question) bool
}
