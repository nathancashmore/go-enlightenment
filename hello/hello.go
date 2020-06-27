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

Provides a greeting in a number of languages such as:
[ "English", "French", "Spanish", "German" ]
*/
package hello

const spanish = "Spanish"
const french = "French"
const german = "German"

const defaultGreeting = "World"
const englishGreeting = "Hello "
const spanishGreeting = "Hola "
const frenchGreeting = "Bonjour "
const germanGreeting = "Hallo "

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
	case german:
		greeting = germanGreeting
	default:
		greeting = englishGreeting
	}
	return
}
