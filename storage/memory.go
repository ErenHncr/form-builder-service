package storage

import (
	"fmt"

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
		Label: types.QuestionLabel{
			"EN": "Read and Approve",
			"TR": "Okudum onaylıyorum",
		},
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
		Label: types.QuestionLabel{
			"EN": "Read and Approve",
			"TR": "Okudum onaylıyorum",
		},
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

func (s *MemoryStorage) CreateQuestion(question types.Question) error {
	questions = append(questions, question)

	return nil
}

func (s *MemoryStorage) DeleteQuestion(id string) error {
	questionIndex := -1
	for i, element := range questions {
		if element.ID == id {
			questionIndex = i
			break
		}
	}

	if questionIndex == -1 {
		return fmt.Errorf("invalid_id")
	}

	questionsCopy := make([]types.Question, 0)
	questionsCopy = append(questionsCopy, questions[:questionIndex]...)
	questions = append(questionsCopy, questions[questionIndex+1:]...)

	return nil
}
