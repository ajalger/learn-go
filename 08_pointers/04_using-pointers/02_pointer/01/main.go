package main

import "fmt"

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}

func zero(z *int) {
	*z = 0
}
