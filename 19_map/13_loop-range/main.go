package main

import (
	"github.com/fatih/color"
)

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	d := color.New(color.FgCyan)
	c := color.New(color.FgGreen)

	for key, val := range myGreeting {
		d.Println(key, " - ", val)
	}

	for i := 0; i < len(myGreeting); i++ {
		c.Println(i, "-", myGreeting[i])
	}
}
