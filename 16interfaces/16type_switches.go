package main

import (
	"fmt"
)

func main() {

	// A type switch is a construct that permits several type assertions in series.

	var i interface{} = 2.3

	switch a := i.(type) {
	case string:
		fmt.Println("Type is string", a)
	case int:
		fmt.Println("Type is int", a)
	default:
		fmt.Println("Unknown type", a)
	}
}
