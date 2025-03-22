package main

import (
	"fmt"
	"strings"
)

const (
	spanish               = "spanish"
	french                = "french"
	portuguese            = "portuguese"
	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Holla, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Olá, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch strings.ToLower(language) {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
