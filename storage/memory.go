package storage

import (
	"context"
	"encoding/json"
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

func findIndexByID(id string) int {
	questionIndex := -1
	for i, element := range questions {
		if element.ID == id {
			questionIndex = i
			break
		}
	}
	return questionIndex
}

func (s *MemoryStorage) Connect(context.Context) error {
	return nil
}

func (s *MemoryStorage) Disconnect(ctx context.Context) error {
	return nil
}

func (s *MemoryStorage) GetQuestions(pagination types.Pagination, sorting []types.Sorting) ([]types.Question, int, error) {
	filteredQuestions := []types.Question{}

	startingIndex := ((pagination.Page - 1) * pagination.Size)
	if startingIndex < len(questions) {
		for i := startingIndex; i < min(pagination.Size, len(questions)); i++ {
			filteredQuestions = append(filteredQuestions, questions[i])
		}
	}

	return filteredQuestions, len(questions), nil
}

func (s *MemoryStorage) GetQuestion(id string) (*types.Question, error) {
	questionIndex := findIndexByID(id)

	if questionIndex == -1 {
		return nil, fmt.Errorf("invalid_id")
	}

	question := questions[questionIndex]

	return &question, nil
}

func (s *MemoryStorage) CreateQuestion(question types.Question) (*types.Question, error) {
	question.ID = uuid.New().String()
	questions = append(questions, question)

	return &question, nil
}

func (s *MemoryStorage) UpdateQuestion(id string, q types.QuestionPatch) (*types.Question, error) {
	selectedQuestion, err := s.GetQuestion(id)

	if err != nil {
		return nil, err
	}

	qCopy, err := json.Marshal(q)
	if err != nil {
		return nil, fmt.Errorf("invalid_marshal_operation")
	}

	id = selectedQuestion.ID
	json.Unmarshal(qCopy, &selectedQuestion)
	selectedQuestion.ID = id

	return selectedQuestion, nil
}

func (s *MemoryStorage) DeleteQuestion(id string) error {
	questionIndex := findIndexByID(id)

	if questionIndex == -1 {
		return fmt.Errorf("invalid_id")
	}

	questionsCopy := make([]types.Question, 0)
	questionsCopy = append(questionsCopy, questions[:questionIndex]...)
	questions = append(questionsCopy, questions[questionIndex+1:]...)

	return nil
}
