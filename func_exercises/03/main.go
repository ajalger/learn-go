package main

import "fmt"

func main() {
	g := greatestNum(23, 32, 56, 12, 100, 45, 26, 2)
	fmt.Println(g)
}

func greatestNum(nums ...float64) float64 {
	var greatest float64
	for _, v := range nums {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}