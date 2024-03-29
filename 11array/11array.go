package main

import (
	"fmt"
	"reflect"
)

func main() {
	// You can initialize an array with pre-defined values using an array literal. [4]int{1, 2, 3, 4} is array literal
	// An array holds a specific number of elements, and it cannot grow or shrink

	arr := [5]int{1, 2, 3} // array
	// When an array declare using an array literal, values can be initialize for specific elements.
	var arr2 [4]int = [4]int{2: 1, 2} // array
	// var arr2 [4]int = [4]int{2: 1, 2, 3} error because out of range
	// 2 is indexed at 3 and 3 is indexed at 4
	var arr3 []int = []int{1, 2, 3, 4} // slice
	arr4 := [...]int{1, 2, 3, 4} // array
	arr7 := []int{1, 2, 3} // slice

	// Copy Array
	arr5 := arr4  // data is passed by value
	arr6 := &arr4 // data is passed by reference

	fmt.Println(reflect.ValueOf(arr).Kind())
	fmt.Println(reflect.ValueOf(arr2).Kind())
	fmt.Println(reflect.ValueOf(arr3).Kind())
	fmt.Println(reflect.ValueOf(arr4).Kind())
	fmt.Println(reflect.ValueOf(arr5).Kind())
	fmt.Println(reflect.ValueOf(arr6).Kind())
	fmt.Println(reflect.ValueOf(arr7).Kind())

	fmt.Printf("%T, %v\n", arr, arr)
	fmt.Printf("%T, %v\n", arr2, arr2)
	fmt.Printf("%T, %v\n", arr3, arr3)
	fmt.Printf("%T, %v\n", arr4, arr4)
	fmt.Printf("%T, %v\n", arr5, arr5)
	fmt.Printf("%T, %v\n", arr6, arr6)
	fmt.Printf("%T, %v\n", arr7, arr7)

}
