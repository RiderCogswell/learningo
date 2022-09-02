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

	if language == spanish {
		return spanishPrefix + name
	} else if language == french {
		return frenchPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Rider", ""))
}