package main

const spanish = "Spanish"

const defaultGreeting = "World"
const englishGreeting = "Hello "
const spanishGreeting = "Hola "

func Hello(name string, lang string) string {
	if name == "" {
		name = defaultGreeting
	}
	return greetingWithPrefix(lang) + name
}

func greetingWithPrefix(lang string) (greeting string) {
	switch lang {
	case spanish:
		greeting = spanishGreeting
	default:
		greeting = englishGreeting
	}
	return
}
