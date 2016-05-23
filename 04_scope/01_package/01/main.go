package main

import "fmt"

var a int = 434

func main() {
	// order matters - block level scope
	fmt.Printf("%v\n", a)
	foo()
	// y is only scoped to this block
	y := 873
	fmt.Printf("%v\n", y)
}

func foo() {
	fmt.Printf("%v from foo\n", a)
}
