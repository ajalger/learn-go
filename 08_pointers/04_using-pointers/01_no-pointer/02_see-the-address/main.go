package main

import "fmt"

func main() {
	x := 5
	fmt.Printf("%p printf main\n", &x) // print memory address
	fmt.Println("println main", &x)
	zero(x)
	fmt.Println(x)
}

func zero(z int) {
	fmt.Printf("%p printf zero \n", &z)
	fmt.Println("println zero", &z)
	z = 0
}