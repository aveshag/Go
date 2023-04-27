package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// slice -> pointer (start location), length, capacity
	primes := [6]int{2, 3, 5, 7, 11, 13}

	chunk := primes[1:3]

	fmt.Println(primes)
	fmt.Println(chunk)
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	fmt.Println(len(chunk))
	fmt.Println(cap(chunk))
	fmt.Println(reflect.ValueOf(chunk).Kind())
	fmt.Println(&primes, &chunk)

	chunk[1] = 45
	fmt.Println(primes)
	fmt.Println(chunk)

	// []bool{true, true, false} is a slice literal
	// declaration and initialization
	chunk = []int{1, 2, 3, 4, 5, 6, 7}
	// low <= high and high <= len(input)
	fmt.Println(chunk)
	fmt.Println("chunk: ", chunk[len(chunk)-1:])
	fmt.Println("chunk: ", chunk[len(chunk):])
	fmt.Println("chunk: ", chunk[len(chunk):len(chunk)])
	// fmt.Println("chunk: ", chunk[len(chunk) + 1:]) // error

	var s []int
	fmt.Println(s, s[:], len(s), cap(s)) // s[0] error
	if s == nil {
		fmt.Println("nil!")
	}

	b := make([]int, 2, 5) // len(b)=2, cap(b)=5 capacity is optional
	fmt.Println(b)

	// slice of slice

	// Below declaration also works
	// grid := [][]string{
	// 	[]string{"_", "|", "_", "|", "_"},
	// 	[]string{"_", "|", "_", "|", "_"},
	// 	[]string{" ", "|", " ", "|", " "},
	// }
	grid := [][]string{
		{"_", "|", "_", "|", "_"},
		{"_", "|", "_", "|", "_"},
		{" ", "|", " ", "|", " "},
	}

	for i := 0; i < len(grid); i++ {
		fmt.Println(strings.Join(grid[i], ""))
	}

	fmt.Println(b, len(b), cap(b))
	c := append(b, 1, 2, 3) // append(b, 1, 2, 3, 4) check capacity of b
	d := append(b, 4, 5, 6)
	e := append(b, 8, 9, 10, 11)
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c)) // value of slice c also changed
	fmt.Println(d, len(d), cap(d))
	fmt.Println(e, len(e), cap(e)) // e doesn't change value of c and d as underlying array changed

	// The range form of the for loop iterates over a slice or map.
	// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

	for i, v := range b {
		fmt.Println(i, v)
	}

	fmt.Println(append(b, c...))

	x := [][]uint8{}
	println(x)

	// Full slice expression has the following syntax:
	// input[low:high:max]
	// Indices low and high work in the same way as with simple slice expressions. The only difference is max which sets result’s capacity to (max - low)

	// y := []int is a error

	// func copy(dst, src []type) int (must be of same datatype)
}
