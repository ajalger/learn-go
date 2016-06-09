package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)
	//delete(myGreeting, "two")
	//delete(myGreeting, "three")
	myGreeting["four"] = "What Up!"
	myGreeting["one"] = "Asalamu Alaykum"
	myGreeting["five"] = "Domo Arigato"

	fmt.Println(len(myGreeting))

	delete(myGreeting, "zero")
	delete(myGreeting, "one")
	delete(myGreeting, "two")

	fmt.Println(len(myGreeting))

	fmt.Println(myGreeting)
}
