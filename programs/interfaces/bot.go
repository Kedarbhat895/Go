package main

import "fmt"

type bot interface {
	getGreeting() (string, string)
}

type englishBot struct{}

type konkaniBot struct{}

func main() {

	eb := englishBot{}
	kb := konkaniBot{}

	printGreeting(eb)
	printGreeting(kb)

}

func printGreeting(b bot) {

	fmt.Println(b.getGreeting())
}
func (englishBot) getGreeting() (string, string) {

	return "hello", "how are you"
}

func (konkaniBot) getGreeting() (string, string) {

	return "namaskaru", "kashi assa"
}
