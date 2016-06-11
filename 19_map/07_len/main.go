package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"
	myGreeting["Henry"] = "Woof"

	fmt.Println(len(myGreeting))
}
