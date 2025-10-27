package main

import "fmt"

type Verte struct {
	Lat, Long float64
}

var m map[string]Verte
// map: diccionario clave→valor sin orden; claves únicas, requiere make() para inicializar.



func main6() {
	m = make(map[string]Verte)
	m["Bell Labs"] = Verte{
		40.68433, -74.39967,
	}
	m["Valen"] = Verte{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Valen"])

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}


// Map literal: crea el map y lo inicializa con 2 entradas.
// Nota: en un map literal las **claves** son obligatorias.
/*var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},     // Vertex{...} puede abreviarse a {...}
	"Google":    {37.42202, -122.08408},
}*/