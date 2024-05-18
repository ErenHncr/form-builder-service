package storage

import (
	"github.com/erenhncr/go-api-structure/types"
	"github.com/google/uuid"
)

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

var questions = []types.Question{
	{
		ID:         uuid.New().String(),
		Key:        "approveConsentAndInformationalText",
		Type:       "checkbox",
		IsRequired: true,
		Label:      "I have read and approved the terms regarding the information and documents to be provided through the platform within the scope of the Law on the Protection of Personal Data No. 6698",
		Links: []types.QuestionLink{{
			Text: "şartları",
			URL:  "https://google.com",
		}},
	},
}

func addExampleQuestion() error {
	exampleQuestion := types.Question{
		ID:         uuid.New().String(),
		Key:        "approveConsentAndInformationalText",
		Type:       "checkbox",
		IsRequired: true,
		Label:      "I have read and approved the terms regarding the information and documents to be provided through the platform within the scope of the Law on the Protection of Personal Data No. 6698",
		Links: []types.QuestionLink{{
			Text: "şartları",
			URL:  "https://google.com",
		}},
	}
	questions = append(questions, exampleQuestion)
	return nil
}

func init() {
	addExampleQuestion()
}

func (s *MemoryStorage) GetQuestions(pagination types.Pagination) []types.Question {
	filteredQuestions := []types.Question{}

	startingIndex := ((pagination.Page - 1) * pagination.Size)
	if startingIndex < len(questions) {
		for i := startingIndex; i < min(pagination.Size, len(questions)); i++ {
			filteredQuestions = append(filteredQuestions, questions[i])
		}
	}

	return filteredQuestions
}

func (s *MemoryStorage) AddQuestion(question types.Question) bool {
	questions = append(questions, question)

	return true
}

func (s *MemoryStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}
}
