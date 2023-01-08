package main

import (
	"fmt";
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	x := 2; var ans float64
	if (x>1) {
		ans = math.Sqrt(float64(x))
	}
	// Variables declared by the statement are only in scope until the end of the if.
	if i:= 1; x>1 {
		ans = math.Sqrt(float64(x+i))
	}
	fmt.Println(ans)
	// Both calls to pow return their results before the call to fmt.Println in main begins.
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}