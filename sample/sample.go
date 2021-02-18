package main

import (
	"fmt"
)

type person struct {
	name string
}

func getSarah() *person {
	sarah := &person{name: "Sarah Palmer"}
	return sarah
}

func main() {
	sarah := getSarah()
	fmt.Println(sarah) // this is what I thought you were asking

	// this prints &{Sarah Palmer} as described
	// in %v description of fmt package

	fmt.Println(sarah.name)
}
