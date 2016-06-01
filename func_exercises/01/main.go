package main

import "fmt"

func main() {
	//fmt.Println(half(2))
	//fmt.Println(half(1))
	h, even := half(2)
	fmt.Println(h, even)
	q, e := half(1)
	fmt.Println(q, e)
}

func half(num int) (int, bool) {
	if num % 2 == 0 {
		return num / 2, true
	}
	return 0, false
//
//
//}

func half (num int) (float64, bool) {
	return float64(num) / 2, num % 2 == 0
}