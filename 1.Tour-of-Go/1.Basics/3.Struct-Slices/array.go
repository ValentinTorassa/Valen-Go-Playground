package main

import "fmt"

// Devuelve un slice de int;
// el literal []int crea el array subyacente y el slice que lo referencia.
func arrayIntro() (primes []int) {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes = []int{2, 3, 5, 7, 11, 13} // ahora es un slice para que sea un array deberia ser
	// [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	return
}

func slices(primes []int) {

	var s []int = primes[1:4]
	fmt.Println(s)
}

func main3() {

	slices(arrayIntro())
	names()

	// Array literal: tipo [3]bool, longitud fija = 3
	//arr := [3]bool{true, true, false}

	// Slice literal: tipo []bool; crea un array interno y un slice que lo referencia
	//slc := []bool{true, true, false} // (ptr, len=3, cap=3) → apunta a un array con esos valores

	sliceBounds()

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice1(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice1(s)

	// Extend its length.
	s = s[:4]
	printSlice1(s)

	// Drop its first two values.
	s = s[2:]
	printSlice1(s)

	// A nil slice has a length and capacity of 0 and has no underlying array.
}

// slices are references to arrays
func names() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceBounds() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func printSlice1(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
