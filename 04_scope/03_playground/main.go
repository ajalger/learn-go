package main

import "fmt"

func main() {
	fmt.Print("Please enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 3

	fmt.Println(output)
}
