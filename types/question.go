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

type validationKey string

const (
	invalidKey        = "invalid_key"
	invalidType       = "invalid_type"
	invalidLabel      = "invalid_label"
	invalidLabelKey   = "invalid_label_key"
	invalidLabelValue = "invalid_label_value"
)

var validationError = map[validationKey]string{
	invalidKey:        invalidKey,
	invalidType:       invalidType,
	invalidLabel:      invalidLabel,
	invalidLabelKey:   invalidLabelKey,
	invalidLabelValue: invalidLabelValue,
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

type QuestionLabel map[LanguageKey]string

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

func (q *Question) MustHaveKey() string {
	if q.Key == "" {
		return fmt.Errorf(validationError[invalidKey]).Error()
	}
	return ""
}

func (q *Question) MustHaveType() string {
	if QuestionTypeMap[q.Type] == "" {
		return fmt.Errorf(validationError[invalidType]).Error()
	}
	return ""
}

func (q *Question) MustHaveLabel() string {
	message := ""

	if q.Label == nil || len(q.Label) == 0 {
		message = fmt.Errorf(validationError[invalidLabel]).Error()
		return message
	}

	for labelKey := range q.Label {
		if Language[labelKey] == "" {
			message = fmt.Errorf(validationError[invalidLabelKey]).Error()
			return message
		} else if q.Label[labelKey] == "" {
			message = fmt.Errorf(validationError[invalidLabelValue]).Error()
		}
	}

	return message
}

func (q *Question) Validate() (bool, []string) {
	validations := []string{q.MustHaveKey(), q.MustHaveType(), q.MustHaveLabel()}
	isValid := true
	validationErrors := make([]string, 0)

	for _, validation := range validations {
		if validation != "" {
			isValid = isValid && false
			validationErrors = append(validationErrors, validation)
		}
	}

	return isValid, validationErrors
}
