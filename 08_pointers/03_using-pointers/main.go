package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a) // 43
	fmt.Println("this is the memory reference to A", &a)

	var b *int = &a
	fmt.Println("this is a B pointer to the A memory reference", b)

	// assigns the memory address of A to 42 from B's pointer
	*b = 42
	fmt.Println(a) // 42
}