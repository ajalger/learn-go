package main

import (
	"strconv"
	"fmt"
)

func main() {
	p := getLargestProduct()
	fmt.Println(p)
}

func getLargestProduct() int {
	var arr []int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			m := multiplyToString(i, j)
			if palindrome(m) {
				m, _ := strconv.Atoi(m)
				arr = append(arr, m)
			}
		}
	}
	return largestNumber(arr)
}

func largestNumber(arr []int) int {
	var largest int
	for _, v := range arr {
		if largest < v {
			largest = v
		}
	}
	return largest
}

func multiplyToString(num1 int, num2 int) string {
	total := num1 * num2
	return strconv.Itoa(total)
}

func palindrome(num string) bool {
	var length = len(num)
	for i := 0; i < (length / 2); i++ {
		if num[i] != num[length - 1 - i] {
			return false
		}
	}
	return true
}