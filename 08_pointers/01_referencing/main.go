package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println("this is the memory reference to A", &a)

	var b *int = &a
	fmt.Println("this is a B pointer to the A memory reference", b)
}
