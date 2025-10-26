package main

import "fmt"

// fmt: paquete estándar para E/S y formateo (Print*, Sprintf, Fprintf, etc.)

func Hello(name string) string {
	return fmt.Sprintf("Hola soy %s", name)
	// Sprintf: arma y devuelve el string sin imprimir
}
