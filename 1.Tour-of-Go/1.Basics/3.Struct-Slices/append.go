package main

import "fmt"

func main5() {
	var s []int
	printSlice2(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice2(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice2(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice2(s)

	s = append(s, 422, 252)
	printSlice2(s)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	/*
	You can skip the index or value by assigning to _.

	for i, _ := range pow
	for _, value := range pow

	If you only want the index, you can omit the second variable.

	for i := range pow
	*/

}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
