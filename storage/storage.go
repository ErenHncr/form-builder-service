package storage

import "github.com/erenhncr/go-api-structure/types"

type Storage interface {
	GetQuestions(types.Pagination) []types.Question
	AddQuestion(types.Question) error
}
