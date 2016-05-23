package main

import "fmt"

const p string = "hello how are you"

func main() {
	const q int = 1323
	i := 64

	for j := i; j <= 70; j++ {
		i += j
	}

	x := 1

	for x <= 10 {
		fmt.Println(x);
		x++
	}

	fmt.Println(i)
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}

// messing around with constants - can't add a constant inside a for loop