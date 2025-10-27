// Una variable de tipo interface guarda un PAR: (valor, tipo-concreto).
// Al llamar un método, Go despacha al método del TIPO CONCRETO guardado.

package main

import (
	"fmt"
	"math"
)

type Abser interface{ Abs() float64 }

type MyFloat float64
func (f MyFloat) Abs() float64 { if f < 0 { return float64(-f) }; return float64(f) }

type Vertex struct{ X, Y float64 }
func (v *Vertex) Abs() float64 { return math.Hypot(v.X, v.Y) }

func main() {
	var a Abser        // (nil, nil): interface cero
	fmt.Println(a == nil) // true

	a = MyFloat(-3)    // (valor=-3, tipo=MyFloat)
	fmt.Printf("%T -> %v\n", a, a.Abs()) // MyFloat -> 3

	v := &Vertex{3, 4}
	a = v              // (valor=&{3 4}, tipo=*Vertex)
	fmt.Printf("%T -> %v\n", a, a.Abs()) // *main.Vertex -> 5

	// Nota: si a guarda un puntero nil de tipo *Vertex, el "valor" no es nil-interface;
	// es (nil, *Vertex). Llamar métodos puede panickear si el método no maneja nil.
	var pv *Vertex = nil
	a = pv             // (nil, *Vertex) → a != nil
	fmt.Println("a es nil?", a == nil) // false
}
