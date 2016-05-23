package main

import (
	"fmt"
	"github.com/ajalger/04_scope/01_package/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	fmt.Println(vis.Different("This is the different function"))
	vis.Print()
}