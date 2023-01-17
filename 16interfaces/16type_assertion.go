package main

import (
	"fmt"
)

func main() {
	// A type assertion provides access to an interface value's underlying concrete value.

	var i interface{} = "hello"

	a := i.(string)
	fmt.Println(a)

	// b := i.(int) // panic: interface conversion: interface {} is string, not int
	b, ok := i.(int)
	fmt.Println(ok, b)
}
