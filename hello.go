package main

import "fmt"

const greeting = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return greeting + name
}

func main() {
	fmt.Printf(Hello("World"))
}
