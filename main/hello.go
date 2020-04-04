package main

import (
	"fmt"
	"os"
)

const spanish = "Spanish"
const french = "French"

const defaultGreeting = "World"
const englishGreeting = "Hello "
const spanishGreeting = "Hola "
const frenchGreeting = "Bonjour "

func main() {
	if len(os.Args[1:]) < 2 {
		fmt.Println("Usage: ./hello <name> <language>")
		os.Exit(9)
	}

	name := os.Args[1]
	lang := os.Args[2]

	fmt.Println(Hello(name, lang))
}

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
