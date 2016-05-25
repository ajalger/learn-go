package main

import "fmt"

func main() {

	var smallNum, bigNum int

	fmt.Println("Please enter a small number: ")
	fmt.Scan(&smallNum)
	fmt.Println("Please enter a larger number: ")
	fmt.Scan(&bigNum)

	remainder := bigNum % smallNum
	fmt.Printf("Your remainder is %v\n", remainder)
}
