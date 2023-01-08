package main

import (
	"fmt";
	"runtime";
	"time"
)

func main() {
	// break statement automatically provided after each case in Go. 
	// Go's switch cases need not be constants, and the values involved need not be integers.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Darwin")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Other")
	}

	switch 1 {
	case 2 - 1:
		fmt.Println("2 - 1")
	case 2 - 2:
		fmt.Println("2 - 2")
	// Error if uncomment below lines
	// case 3 - 2:
	// 	fmt.Println("3 - 2")
	}

	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}