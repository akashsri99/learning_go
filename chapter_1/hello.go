package main

const englishPrefix = "Hello, "

const (
	englishLanguage string = "English"
	spanishLanguage string = "Spanish"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix
	if language == "Spanish" {
		prefix = "Hola, "
	}
	return prefix + name
}

func getPrefixByLanguage(language string) (prefix string) {

	switch language {
	case englishLanguage:
		prefix = "Hello, "
		break
	case spanishLanguage:
		prefix = "Hola, "
		break
	}
	return prefix
}

func main() {
	// fmt.Println(Hello())
}
