package storage

import "github.com/erenhncr/go-api-structure/types"

type Storage interface {
	GetQuestions(types.Pagination) []types.Question
	GetQuestion(string) (*types.Question, error)
	CreateQuestion(types.Question) error
	DeleteQuestion(string) error
	UpdateQuestion(string, types.Question) (*types.Question, error)
}
