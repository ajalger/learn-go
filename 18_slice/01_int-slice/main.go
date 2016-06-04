package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	newSlice := make([]int, 50, 100)
	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
}
