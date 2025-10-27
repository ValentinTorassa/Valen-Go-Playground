package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	v4 = &Vertex{1, 2} // has type *Vertex
)

func main2() {
	i, j := 42, 2701

	//  The & operator generates a pointer to its operand.
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	//  The * operator denotes the pointer's underlying value.
	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	pa := &v
	pa.X = 1e9
	fmt.Println(v)
	fmt.Println(v1, v4, v2, v3)
}
