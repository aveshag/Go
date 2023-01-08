package main

import "fmt"

// func swap(p int, q int) (int, int) {
// 	var x = q
// 	var y = p
// 	return x, y
// }

// Go's return values may be named. If so, they are treated as 
// variables defined at the top of the function.
// A return statement without arguments returns the named return values.
// This is known as a "naked" return.
// Naked return statements should be used only in short functions, 
// as with the example shown here. They can harm readability in longer functions.

func swap(p int, q int) (x, y int) {
	x = q
	y = p
	return
}

func main() {
	fmt.Println(swap(3, 5))
}