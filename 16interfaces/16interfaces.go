package main

import (
	"fmt"
	"math"
)

type edge struct {
	a, b int
}

type I interface {
	dist() float64
}

// A type implements an interface by implementing its methods
// This method means type edge implements the interface I,
// but we don't need to explicitly declare that it does so.
func (e *edge) dist() float64 {
	if e != nil {
		return math.Sqrt(float64(e.a*e.a + e.b*e.b))
	} else {
		return 0.0
	}
}

func (e edge) myabs() float64 {
	return math.Abs(float64(e.a - e.b))
}

func main() {
	var i I = &edge{1, 2}
	// Under the hood, interface values can be thought of as a tuple of a value and a concrete type: (value, type)
	// An interface value holds a value of a specific underlying concrete type.
	// Calling a method on an interface value executes the method of the same name on its underlying type.
	describe(i)
	fmt.Println(i.dist())
	// fmt.Println(i.myabs())

	// Interface values with nil underlying values
	// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
	// Note that an interface value that holds a nil concrete value is itself non-nil.
	var e *edge
	i = e
	describe(i)
	i.dist()

	// Nil interface values
	// A nil interface value holds neither value nor concrete type.
	// Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	// var i2 I
	// describe(i2)
	// i2.dist()


	// The empty interface
	// The interface type that specifies zero methods is known as the empty interface: interface{}
	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}

	var i3 interface{}
	describe2(i3)

	i3 = 42
	describe2(i3)

	i3 = "hello"
	describe2(i3)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
