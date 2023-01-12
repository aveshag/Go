package main

import "fmt"

type vertex struct {
	a, b float32
}

// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
func (v vertex) difference() float32 {
	return v.b - v.a	
}

// Regular function
func difference2(v vertex) float32 {
	return v.b - v.a
}

// You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

type Myint int

func (i Myint) square() int {
	return int(i)*int(i)
}

// pointer receiver

func (v *vertex) scale() {
	v.a = v.a + 2
	v.b = v.b + 3
}

func main() {
	ver := vertex{1, 3}

	// p := &ver 
	// the method call p.difference() is interpreted as (*p).difference()
	// fmt.Println(p.difference())

	fmt.Println(ver.difference())

	// methods with pointer receivers take either a value or a pointer as the receiver when they are called
	// p := &ver
	// p.scale() // OK
	// Go interprets ver.scale() as (&ver).scale()
	ver.scale() // OK

	fmt.Println(difference2(ver))
	var mi Myint = 5
	mi2 := Myint(5)

	fmt.Println(mi.square())
	fmt.Println(mi2.square())

	// There are two reasons to use a pointer receiver.
	// The first is so that the method can modify the value that its receiver points to.
	// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct,

}
