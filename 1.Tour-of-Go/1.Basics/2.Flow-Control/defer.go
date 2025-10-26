package main

import "fmt"

func main() {
	defer fmt.Println("world") // Se ejecuta AL FINAL de main 

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// Los argumentos se EVALÚAN acá (guarda el valor actual de i),
		// pero la ejecución se difiere. Como el stack de defers es LIFO,
		// al final se imprimirá: 9,8,7,...,0.
		defer fmt.Println(i)
	}

	defer fmt.Println("TEST") // se imprime antes que los otros dos defers

	//defer usa LIFO Last In First Out

	fmt.Println("done")
}
