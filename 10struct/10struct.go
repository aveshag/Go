package main

import "fmt"

func main() {
	type vertex struct {
		x int // no var
		y int // no var
		z, w int
	}

	v := vertex{1, 2, 9, 10}
	fmt.Println(v)
	v.x = 5
	v.y, v.x = 6, 7
	fmt.Println(v)
	v.y, v.x = v.x, v.y
	fmt.Println(v)

	var p *vertex
	p = &v
	fmt.Println(&p)
	fmt.Println(&v)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(p.x)
	fmt.Println((*p).x)

	var (
		// v1 = vertex{1, 2}  // has type Vertex // too few values in vertex
		v1 = vertex{1, 2, 3, 4}
		v2 = vertex{x: 1}  // y:0, z:0 and w:0  are implicit
		v3 = vertex{}      // x:0, y:0, z:0 and w:0
	)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	
}