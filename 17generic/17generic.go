package main

import "fmt"

// Go functions can be written to work on multiple types using type parameters. 
// The type parameters of a function appear between brackets, before the function's arguments
func item[T comparable] (s []T, x T) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	sl := []int{1,2,3,4,5}
	fmt.Println(item(sl, 3))

	sl2 := []string{"abc", "xyz", "pqr"}
	fmt.Println(item(sl2, "ab"))

	y := List[int]{nil, 5}
	x := List[int]{&y, 2}
	fmt.Println(x)
	fmt.Println(*x.next)
}