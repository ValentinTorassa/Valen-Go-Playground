package main

import "fmt"

// Sqrt calcula la raíz cuadrada de x usando Newton-Raphson.
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z) // mostrar el progreso de z en cada iteración
	}
	return z // z queda muy cerca de √x
}

func main4() {
	fmt.Println(Sqrt(2))
}
