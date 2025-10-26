package main

import "fmt"

func main2() {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += i
		fmt.Println(sum)
	}

	nombre := "Valen"
	fmt.Printf("La suma final es %d, %s\n", sum, nombre)

	// while-style
	sum2 := 1
	for sum2 < 1000 { //or  for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
	/*forever
	for {
	}*/
}
