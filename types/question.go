package types

import (
	"fmt"

	"github.com/google/uuid"
)

type QuestionType string

const (
	QuestionTypeText     = "TEXT"
	QuestionTypeCheckbox = "CHECKBOX"
)

var QuestionTypeMap = map[QuestionType]string{
	QuestionTypeText:     QuestionTypeText,
	QuestionTypeCheckbox: QuestionTypeCheckbox,
}

func (t QuestionType) String() string {
	return QuestionTypeMap[t]
}

type QuestionLink struct {
	Text         string `json:"text"`
	URL          string `json:"url"`
	DocumentType string `json:"documentType"`
}

type QuestionNote struct {
	Text   string `json:"text"`
	Detail string `json:"detail"`
}

type QuestionOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type QuestionParent struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

type QuestionVideo struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	URL      string `json:"url"`
}

type QuestionDisplay struct {
	Mode       string                     `json:"mode"`
	Conditions []QuestionDisplayCondition `json:"conditions"`
}

type QuestionDisplayCondition struct {
	Condition string      `json:"condition"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
	// accepts number, string string array number array
}

type QuestionLabel map[Language]string

type Question struct {
	ID         string           `json:"id"`
	IsRequired bool             `json:"isRequired"`
	Key        string           `json:"key"`
	Label      QuestionLabel    `json:"label"`
	Links      []QuestionLink   `json:"links"`
	Type       QuestionType     `json:"type"`
	Notes      []QuestionNote   `json:"notes"`
	Options    []QuestionOption `json:"options"`
	Parent     *QuestionParent  `json:"parent"`
	PDF        string           `json:"pdf"`
	Video      *QuestionVideo   `json:"video"`
	Display    *QuestionDisplay `json:"display"`
}

func NewQuestion() *Question {
	return &Question{
		ID:         uuid.New().String(),
		IsRequired: false,
		Type:       QuestionTypeText,
		Links:      []QuestionLink{},
		Notes:      []QuestionNote{},
		Options:    []QuestionOption{},
	}
}

func (q *Question) MustHaveType() string {
	if QuestionTypeMap[q.Type] == "" {
		return fmt.Errorf("must have a valid type").Error()
	}
	return ""
}

func (q *Question) MustHaveLabel() string {
	message := ""

	if q.Label == nil || len(q.Label) == 0 {
		message = fmt.Errorf("must have a label").Error()
		return message
	}

	for labelKey := range q.Label {
		if LanguageMap[labelKey] == "" {
			message = fmt.Errorf("must have a valid label key").Error()
			return message
		} else if q.Label[labelKey] == "" {
			message = fmt.Errorf("must have a valid label value").Error()
		}
	}

	return message
}

func (q *Question) Validate() (bool, []string) {
	var validations = []string{q.MustHaveType(), q.MustHaveLabel()}

	isValid := true
	for _, validation := range validations {
		if validation != "" {
			isValid = false
			break
		}
	}

	return isValid, validations
}
