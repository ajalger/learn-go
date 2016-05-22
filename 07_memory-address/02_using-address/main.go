package main

import "fmt"

const metersToYars float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters travelled: ")

	// use the memory adress to story meters input
	fmt.Scan(&meters)
	yards := meters * metersToYars
	fmt.Println(meters, "meters is", yards, "yards.")
}
