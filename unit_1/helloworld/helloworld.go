package main

const (
	greetingStarting = "Hello "
	frenchGreeting   = "Bonjure "
	frenchLanguage   = "French"
	spanishLanguage  = "Spanish"
	spanishGreeting  = "Hola "
	hindiGreeting    = "Namaste "
	hindiLanguage    = "Hindi"
)

func main() {

}

func greetingMessage(name string, language string) string {

	if name == "" {
		return "Hello Guys"
	}
	return greetingText(language) + name
}

func greetingText(language string) (prefix string) {
	switch language {
	case frenchLanguage:
		prefix = frenchGreeting
	case spanishLanguage:
		prefix = spanishGreeting
	case hindiLanguage:
		prefix = hindiGreeting
	default:
		prefix = greetingStarting
	}
	return
}
