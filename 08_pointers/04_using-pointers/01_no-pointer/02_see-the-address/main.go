package main

import "fmt"

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // print memory address
	fmt.Println(&x)
	zero(x)
	fmt.Println(x)
}

func zero(x int) {
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	x = 0
}