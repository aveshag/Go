package main

import (
	"fmt"
)

func main() {
	// var p *int
	v := 5

	p := &v
	// Unlike C, Go has no pointer arithmetic.
	fmt.Println(p)
	fmt.Println(*p + 1)
	fmt.Printf("%T\n", p)
}