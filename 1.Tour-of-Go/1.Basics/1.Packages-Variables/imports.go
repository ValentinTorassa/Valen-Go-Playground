package main

import (
	"fmt"
	"math"
	"math/rand"
)

var c, python, java bool

var j, f int = 20, 21

// k := 3 Esto no funciona

func main2() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("world", "hello")
	fmt.Println(a, b)
	fmt.Println(split(20))
	var i int
	i = 10
	fmt.Println(i, c, python, java)
	fmt.Println(j, f)
	k := 3
	fmt.Println(k)
}

// func add(x, y int) cuando dos argumentos comparten un tipo
func add(x int, y int) int {
	return x + y
}

// una funcion puede retornar varios resultados
func swap(x, y string) (string, string) {
	return y, x
}

// si ya defini que nombre voy a devolver puedo devolver un return naked
func split(sum int) (x, y int) {
	x = sum / 5
	y = sum - x
	return
}
