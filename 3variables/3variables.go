package main

import (
	"fmt"
	"math/cmplx"
)

// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32
// represents a Unicode code point
// float32 float64
// complex64 complex128

var (
	MaxInt uint32     = 1<<32 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var str = "golang"
	fmt.Println(str)
	var a int
	fmt.Println(a)
	var b bool = true
	fmt.Println(b)
	var c, d int = 3, 5
	// var c = 3, d int = 3 Error
	// var c = 3, d = 3 Error
	// var c, d = 3, 5
	// var c, d int = 3, 5, aa = 3 Error
	// var c int = 3, d bool = 5
	fmt.Println(c, d)
	abc := 10
	fmt.Println(abc)

	// %v: value in a default format
	// %d: decimal integer
	// %T: type of value
	// %c: character
	// %q: quoted character/string
	// %s: plain string
	// %t: true or false
	// %f: floating numbers
	// %.2f: floating numbers upto 2 decimal places
	fmt.Printf("Type = %T, Value = %v\n", z, z)
	fmt.Printf("Type = %T, Value = %v\n", 12, 12)
	// Unlike in C, in Go assignment between items of different type requires an explicit conversion.

	var t1 int = 1
	var t2 float32 = float32(t1)

	fmt.Println(t1, t2)

	// Constants cannot be declared using the := syntax.
	// untyped const
	const clx = 2 + 3i
	clx = 2
	fmt.Println(clx)
	// typed const
	const name string = "abc"
}
