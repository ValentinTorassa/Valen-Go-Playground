package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
	v := 0.867 + 0.5i // change me!
	fmt.Printf("v is of type %T\n", v)
	// no se puede Pi = 3.15
	fmt.Println(Pi)
}