package main

import "fmt"

func main() {
	half := func (num int) {
		if num % 2 == 0 {
			fmt.Println(num / 2, true)
		} else {
			fmt.Println(0, false)
		}
	}
	half(2)
	half(1)
}


