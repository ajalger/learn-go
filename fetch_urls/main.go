package main

import "fmt"

func main() {
	multiples()
	fmt.Println("... done")
}

func multiples() {
	length := 100
	for i := 1; i < length; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Printf("%v is a multiple of 3 and 5\n", i)
		}
	}
}

