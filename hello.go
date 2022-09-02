package main

import "fmt"

const englishHelloPrefix = "hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Rider", ""))
}