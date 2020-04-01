package main

const spanish = "Spanish"
const french = "French"

const defaultGreeting = "World"
const englishGreeting = "Hello "
const spanishGreeting = "Hola "
const frenchGreeting = "Bonjour "

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
	case french:
		greeting = frenchGreeting
	default:
		greeting = englishGreeting
	}
	return
}
