package main

import "fmt"

// Función genérica: T es un parámetro de tipo.
// Constraint: "comparable" → permite usar == y != sobre T.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x { // válido porque T es comparable
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(Index([]int{10, 20, 30}, 20))           // 1
	fmt.Println(Index([]string{"a", "b", "c"}, "b"))    // 1

	// También funciona con tipos definidos por vos si son comparables.
	type ID string
	fmt.Println(Index([]ID{"u1", "u2", "u3"}, ID("u2"))) // 1
}
