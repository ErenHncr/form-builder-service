package types

type LanguageKey string

const (
	LanguageEN = "EN"
	LanguageTR = "TR"
)

var Language = map[LanguageKey]string{
	LanguageEN: LanguageEN,
	LanguageTR: LanguageTR,
}

func (l LanguageKey) String() string {
	return Language[l]
}
