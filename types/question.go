package types

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

type Question struct {
	ID         string           `json:"id"`
	IsRequired bool             `json:"isRequired"`
	Key        string           `json:"key"`
	Label      string           `json:"label"`
	Links      []QuestionLink   `json:"links"`
	Type       string           `json:"type"` // TODO: add enum
	Notes      []QuestionNote   `json:"notes"`
	Options    []QuestionOption `json:"options"`
	Parent     QuestionParent   `json:"parent"`
	PDF        string           `json:"pdf"`
	Video      QuestionVideo    `json:"video"`
	Display    QuestionDisplay  `json:"display"`
}

func (q *Question) validate() bool { return true }
