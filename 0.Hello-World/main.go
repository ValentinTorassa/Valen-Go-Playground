package main

// package main: este folder compila a ejecutable; todos los .go aquí comparten el mismo paquete

import "fmt"

func main() {
	fmt.Println(Hello("Valen"))
}

/*
go mod init ... → crear módulo.
go run . → ejecutar rápido.
go build -o bin/hello . → compilar binario.
gofmt -w . → formatear código.
go test -v → correr tests (si hay).
*/
