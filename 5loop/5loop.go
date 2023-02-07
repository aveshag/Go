package main

import "fmt"

func main() {
	var sum int = 0
	// no parentheses surrounding the three components of the for statement and the braces { } are always required.
	for i := 0; i < 10; i++ {
		sum += i
	}
	// var declaration not allowed in for initializer
	// for var i int = 0; i<10; i++ {
	// 	sum += i
	// }

	// The init and post statements are optional.
	for sum < 50 {
		sum += sum
	}
	for sum < 100 {
		sum += sum
	}
	for sum < 100 {
		sum += sum
	}
	// Infinite loop
	// for ; ; {
	// 	sum += 0
	// }
	// for {
	// }

	fmt.Println(sum)

	for range "Hello" {
		fmt.Println("Hello")
	}
}
