package main

import (
	"fmt"
	"math"
)

func callfunc(fn func(float64, float64) float64) float64 {
	return fn(4, 5)
}

func incrementor() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {

	// Anonymous function

	a := func(a, b float64) float64 {
		return a + b
	}
	fmt.Println(a(1, 10))
	fmt.Println(callfunc(a))
	fmt.Println(callfunc(math.Pow))

	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

	// Closure function
	// Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.
	// closure also provides data isolation.

	fmt.Println("Closure")
	// Let’s take a look at the scope of the incrementor() function (everything inside the curly braces):
	// i := 0
	// return func() int {
	// 	i++
	// 	return i
	// }
	// When the execution of incrementor() finishes, the scope above is destroyed. But the closure is still able to access and update the variable(s) of the scope even though it’s gone.
	next := incrementor()
	next2 := incrementor()
	fmt.Println(next())
	fmt.Println(next2())
	fmt.Println(next())
	fmt.Println(next2())

	j := 0

	counter := func() int {
		j++
		return j
	}
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	j++
	fmt.Println(counter())

}
