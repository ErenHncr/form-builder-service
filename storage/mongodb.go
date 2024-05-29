package storage

import "github.com/erenhncr/go-api-structure/types"

type MongoDBStorage struct{}

func (s *MongoDBStorage) GetQuestions(pagination types.Pagination) []types.Question {
	return []types.Question{}
}

func (s *MongoDBStorage) GetQuestion(id string) (*types.Question, error) {
	return nil, nil
}

func (s *MongoDBStorage) CreateQuestion(question types.Question) error {
	return nil
}
func (s *MongoDBStorage) UpdateQuestion(id string, q types.Question) (*types.Question, error) {
	return nil, nil
}

func (s *MongoDBStorage) DeleteQuestion(id string) error {
	return nil
}
