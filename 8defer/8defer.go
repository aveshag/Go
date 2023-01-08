package main

import (
	"fmt"
)

func main() {
	// Deferred function calls are pushed onto a stack. When a surrounding function returns, its deferred calls are executed in last-in-first-out order.

	fmt.Println("A")
	fmt.Println("B")
	defer fmt.Println("C")
	defer fmt.Println("D")
	fmt.Println("E")
	fmt.Println("F")
}