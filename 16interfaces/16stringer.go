package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	// One of the most ubiquitous interfaces is Stringer defined by the fmt package.
	// type Stringer interface {
	// 	String() string
	// }
	// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Printf("{%T, %v}\n", a, a)
	fmt.Println(a, z)
}