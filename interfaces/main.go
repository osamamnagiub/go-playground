package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Printf("b: %v\n", b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// func (englishBot) printGreeting(eb englishBot) {
// 	fmt.Printf("eb: %v\n", eb.printGreeting())
// }

// func (spanishBot) printGreeting(sb spanishBot) {

// }
