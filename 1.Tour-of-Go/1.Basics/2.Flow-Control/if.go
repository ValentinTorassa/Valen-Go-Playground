package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func mayor(x int) bool {
	if x > 0 {
		return true
	}
	return false
}

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}

func main3() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(mayor(2), mayor(-4))
	fmt.Println(pow(3, 2, 6), pow(3, 3, 20))
}
