package main

import (
	"fmt"
)

func main() {
	// var p *int
	v := 5

	p := &v
	// Unlike C, Go has no pointer arithmetic. then why we have type of pointer? (may be for de-referencing?)
	fmt.Println(p)
	fmt.Println(*p + 1)
	fmt.Printf("%T\n", p)

	// slices and maps are passed by reference by default but not arrays
}