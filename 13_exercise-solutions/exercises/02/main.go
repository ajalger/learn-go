package main

import "fmt"

func main() {
	fmt.Println("Please enter your name: ")
	var input string;
	fmt.Scanf("%s", &input)
	fmt.Println("Hello, my name is " + input)
}
