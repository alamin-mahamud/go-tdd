package locale

const (
	helloPrefixEnglish = "Hello"
	helloPrefixSpanish = "Hola"
	separator          = ", "
	ENGLISH            = "English"
	SPANISH            = "Spanish"
)

var m map[string]string = make(map[string]string)

func initPrefix() {
	m[ENGLISH] = helloPrefixEnglish
	m[SPANISH] = helloPrefixSpanish
}

func getPrefixFromLanguage(language string) string {

	if language == "" {
		language = ENGLISH
	}

	return m[language] + separator
}

func isLanguageAvailable(language string) bool {

	initPrefix()

	if _, ok := m[language]; !ok {
		return false
	}
	return true
}

func Hello(s string, language string) string {

	if s == "" {
		s = "World"
	}

	if language != "" && !isLanguageAvailable(language) {
		return "Unknown Language. Try English"
	}

	return getPrefixFromLanguage(language) + s
}
