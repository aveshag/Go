package main

import (
	"fmt"
	"encoding/json"
)

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
	
	v_p := new(vertex)
	v_p.x = 2
	fmt.Println(v_p)
	fmt.Println(*v_p)

	var v_p2 = new(vertex)
	v_p2.y = 5
	fmt.Println(v_p2)

	// Use Field Tags in the Definition of Struct Type
	// The tags are represented as raw string values (wrapped within a pair of ``) and ignored by normal code execution.

	type Employee struct {
		FirstName  string `json:"firstname"`
		LastName   string `json:"lastname"`
		City string `json:"city"`
	}

	json_string := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`
 
    emp1 := new(Employee)
    json.Unmarshal([]byte(json_string), emp1)
    fmt.Println(emp1)
 
    emp2 := new(Employee)
    emp2.FirstName = "Ramesh"
    emp2.LastName = "Soni"
    emp2.City = "Mumbai"
    jsonStr, _ := json.Marshal(emp2)
    fmt.Printf("%s\n", jsonStr)

}