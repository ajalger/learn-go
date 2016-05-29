package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := 76

	fmt.Println(array)
	fmt.Printf("%T \n", array)
	fmt.Printf("%T \n", b)
	fmt.Println(&b)

	for _, v := range array {
		fmt.Println(v * 10)
	}
}
