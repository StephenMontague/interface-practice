package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// For these functions, we can omit the value (eb/sb) because it is not being used
// EX: func (eb englishBot) --> eb not needed in this case
func (englishBot) getGreeting() string {
	// Very custom logic for generatig an English greeting :)
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating a Spanish greeting :)
	return "Hola!"
}

