package types

type Language string

const (
	LanguageEN = "EN"
	LanguageTR = "TR"
)

var LanguageMap = map[Language]string{
	LanguageEN: LanguageEN,
	LanguageTR: LanguageTR,
}

func (language Language) String() string {
	return LanguageMap[language]
}
