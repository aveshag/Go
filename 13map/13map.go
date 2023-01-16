package main

import "fmt"

type edge struct{
	e1, e2 int
}

var m map[string]edge

func main() {
	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

	if m == nil {
		fmt.Println("Map is nil")
	}
	// Map is nil panic: assignment to entry in nil map
	// m["edge1"] = edge{1, 2}
	// fmt.Println(m)

	m = make(map[string]edge)
	m["edge1"] = edge{1, 2}
	fmt.Println(m)

	// Map literals are like struct literals, but the keys are required.

	// var edges = map[string]edge{
	// 	"edge1": edge{4, 5},
	// 	"edge2": edge{9, 10},
	// } 
	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var edges = map[string]edge{
		"edge1": {4, 5},
		"edge2": {9, 10},
	} 

	fmt.Println(edges)
	// Insert and update
	edges["edge2"] = edge{11, 12}
	edges["edge3"] = edge{-1, -8}
	fmt.Println(edges)

	for i, v := range edges {
		fmt.Println(i, v)
	}

	delete(edges, "edge3")
	fmt.Println(edges)

	// Test that a key is present with a two-value assignment:
	elem, ok := edges["edge3"]
	fmt.Println(elem, ok)

}
