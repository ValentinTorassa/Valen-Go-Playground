package main2

import (
	"fmt"
	"math"
)

type Vertex struct{ X, Y float64 }

// Método con receiver por valor (copia): NO muta el Vertex original.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

//  1. Receiver por VALOR: trabaja sobre una COPIA → NO cambia f afuera.
//     Regla de tu experimento: devuelve |f| + 1 (para notar el cambio).
func (f MyFloat) Abs() float64 {
	if f < 0 {
		f = -f + 1
		return float64(f) // ← NO muta la variable original
	}
	return float64(f + 1)
}

//  2. Receiver por PUNTERO: opera sobre *f → SÍ cambia el valor original.
//     Regla de tu experimento: si es negativo, lo reemplaza por |f| + 1.
func (f *MyFloat) AbsPointer() float64 {
	// f es *MyFloat, por eso usamos *f para leer/escribir el número real.
	if *f < 0 {
		*f = -*f + 1 // ← MUTACIÓN in-place
	}
	return float64(*f)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main2() {
	v := Vertex{10, 4}
	fmt.Println("Vertex.Abs():", v.Abs()) // solo ejemplo de método por valor en otro tipo

	f := MyFloat(-2)
	fmt.Println("start f =", f)                    // -2
	fmt.Println("AbsPointer()  →", f.AbsPointer()) // muta f: pasa a 3 (= |-2|+1)
	fmt.Println("after         =", f)              // 3 (confirmamos la MUTACIÓN)

	fmt.Println("Abs()         →", f.Abs()) // NO muta: devuelve 4 (= |3|+1), f sigue en 3
	fmt.Println("after         =", f)       // 3 (sigue igual)

	// Llamar un método de puntero sobre una variable NO puntero es válido:
	// Go toma &f implícitamente si el valor es addressable.
	fmt.Println("AbsPointer()  →", f.AbsPointer()) // ya es positivo: queda 3
	fmt.Println("after         =", f)              // 3

	g := MyFloat(-5)
	fmt.Println("\nOtro caso g:", g)
	_ = g.Abs()                     // devuelve 6, pero NO muta g
	fmt.Println("g tras Abs():", g) // sigue -5
	_ = (&g).AbsPointer()           // muta g → 6
	fmt.Println("g tras AbsPointer():", g)

	va := Vertex{3, 4}
	Scale(&va, 10)
	// CORRECTO: para una FUNCIÓN que recibe *Vertex, tenés que pasar &va.
	// En cambio, con un MÉTODO de receptor *Vertex, podés llamar v.Scale(...)
	// sin escribir &v: Go toma la dirección implícitamente si el valor es addressable.

	//al revés da error. En funciones “normales” no hay auto-&/auto-*.
	//  Si la firma pide valor (T), no podés pasar un puntero (*T) sin desreferenciar;
	// y si pide puntero, no podés pasar un valor sin tomar su dirección.

	/* There are two reasons to use a pointer receiver.

	The first is so that the method can modify the value that its receiver points to.

	The second is to avoid copying the value on each method call.
	This can be more efficient if the receiver is a large struct, for example. */
}

// En Go, "type" define un **tipo nuevo**. Puede ser:
// 1) Un **struct** (con campos) → tipo compuesto
// 2) Un **tipo definido** basado en otro (nuevo identificador / method set propio)
// 3) Un **alias** de tipo (mismo tipo, otro nombre)

// 1) STRUCT: un agregado de campos con nombre y tipo
/*type Vertex struct {
	X float64
	Y float64
}*/

// 2) Tipo definido (nuevo tipo basado en float64; NO es alias)
//type MyFloat float64

// 3) Alias de tipo (solo otro nombre para el mismo tipo)
//type Rune = int32
