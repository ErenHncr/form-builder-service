package storage

import "github.com/erenhncr/go-api-structure/types"

type Storage interface {
	GetQuestions(types.Pagination) []types.Question
	CreateQuestion(types.Question) error
	DeleteQuestion(string) error
	UpdateQuestion(string, types.Question) (*types.Question, error)
}
