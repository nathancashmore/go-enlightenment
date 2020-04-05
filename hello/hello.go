/*
Hello (name, language) :
Provides examples of the following:
 import - e.g. fmt and os
 func - public & private (uppercase & lowercase)
 const - public & private (uppercase & lowercase)
 if
 switch
 returning declared variable
 processing input arguments
*/
package hello

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

// Provide a greeting for a given name in a number
// of different languages
//
// Supported Languages:
//
//  [ "English", "French", "Spanish" ]
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
