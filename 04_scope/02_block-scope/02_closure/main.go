package main

import "fmt"

func main() {
	increase := incrementer()
	fmt.Println(increase())
	fmt.Println(increase())
}

func incrementer() func() int {
	x := 0
	return func() int {
		x++
		return x;
	}
}

