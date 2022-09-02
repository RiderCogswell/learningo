package main

import "fmt"

const spanish = "Spanish"
const english = "English"
const french = "French"
const englishHelloPrefix = "hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix
	
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Rider", ""))
}