/*
Una interface es un conjunto de firmas de métodos.

Se implementa implícitamente: si un tipo tiene esos métodos, 
ya implementa la interface (sin implements).

Method sets importan:

Para un valor T: solo métodos con receiver T.

Para un puntero *T: métodos con receiver T y *T.

Resultado: si la interface exige un método que solo 
está definido con receiver puntero, solo *T implementa esa interface (no el valor T).
*/


// Interfaces are implemented implicitly:
// No hay "implements": si un tipo tiene los métodos, YA implementa la interface.
// Esto desacopla: la interface puede vivir en un paquete y las implementaciones en otros, sin coordinación previa.


package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

// Receiver por valor: lo implementan MyFloat y *MyFloat
func (f MyFloat) Abs() float64 {
	if f < 0 { return float64(-f) }
	return float64(f)
}

type Vertex struct{ X, Y float64 }

// Receiver por puntero: solo *Vertex implementa Abser (NO Vertex)
func (v *Vertex) Abs() float64 {
	return math.Hypot(v.X, v.Y)
}

func main() {
	var a Abser

	var f MyFloat = -2
	a = f        // OK: MyFloat tiene Abs() por valor
	fmt.Println(a.Abs())

	v := Vertex{3, 4}
	// a = v      // ❌ ERROR: Vertex NO implementa Abser (Abs es sobre *Vertex)
	a = &v       // ✅ OK: *Vertex sí implementa Abser
	fmt.Println(a.Abs())
}

// Comentario: si un método se define con receiver *T, la interface la implementa *T (puntero), no T (valor).
// Si el método se define con receiver T, la implementan T y *T.
