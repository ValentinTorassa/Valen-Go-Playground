// Go Notes — Tipos, Métodos, Punteros, Interfaces, Type Assertions
// Copiá/pegá y ejecutá; cada bloque ilustra un concepto con comentarios cortos.

package main

import (
	"fmt"
	"math"
)

//
// 1) TIPOS en Go (no hay “clases”)
// - type define un tipo nuevo (struct, definido, alias).
// - Podés “pegar” métodos a tus tipos (polimorfismo via interfaces, no herencia).
//

// 1.1) STRUCT: tipo compuesto con campos
type User struct {
	ID   int
	Name string
}

// 1.2) Tipo DEFINIDO (nuevo identificador basado en otro tipo, NO alias)
type Age int

func (a Age) IsAdult() bool { return a >= 18 } // método en tipo definido

// 1.3) Alias de tipo (mismo tipo, otro nombre)
type Rune = int32

//
// 2) MÉTODOS: receiver por VALOR vs por PUNTERO
// - (T) → copia: NO muta el original.
// - (*T) → referencia: PUEDE mutar el original. Go permite v.M() y toma &v implícito para métodos (no para funciones).
//

type Vertex struct{ X, Y float64 }

func (v Vertex) Abs() float64 { // valor: copia
	return math.Hypot(v.X, v.Y)
}
func (v *Vertex) Scale(f float64) { // puntero: muta el original
	v.X *= f
	v.Y *= f
}

// Equivalente con funciones libres (no hay auto &/* en funciones)
func ScaleVal(v Vertex, f float64)  { v.X *= f; v.Y *= f } // no muta al llamante
func ScalePtr(v *Vertex, f float64) { v.X *= f; v.Y *= f } // requiere &v al llamar

//
// 3) Receiver puntero que MODIFICA el valor (demostración simple)
//
type MyFloat float64

func (f MyFloat) AbsVal() float64 { // no muta
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func (f *MyFloat) AbsInPlacePlus1() float64 { // muta: reemplaza por |f|+1
	if *f < 0 {
		*f = -*f + 1
	}
	return float64(*f)
}

//
// 4) INTERFACES: contrato de métodos, implementación IMPLÍCITA
// - Si un tipo tiene los métodos requeridos, ya implementa la interface (no hay “implements”).
// - Method set importa:
//   * T tiene métodos con receiver T.
//   * *T tiene métodos con receiver T y *T.
//   Si la interface requiere un método definido solo en *T, solo *T satisface.
//
type Abser interface{ Abs() float64 }

type MyFloat2 float64

func (f MyFloat2) Abs() float64 { // por valor: T y *T cumplen
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Abs en *Vertex ⇒ solo *Vertex implementa Abser (Vertex no).
func (v *Vertex) Abs() float64 { // (además de Scale) – demuestra method set
	return math.Hypot(v.X, v.Y)
}

//
// 5) Interface values = (valor, tipo-concreto)
// - Una interface guarda un par (value, concrete-type).
// - Llamar un método despacha al método del tipo concreto guardado.
// - Ojo: interface no-nil con valor puntero nil → a != nil
//

//
// 6) TYPE ASSERTIONS (aseveraciones de tipo)
// - t := i.(T) → si i no contiene T → panic.
// - t, ok := i.(T) → ok=false si no es T (sin panic); t es el cero de T.
// - Coincidencia EXACTA (Person vs *Person).
//

// Utilidad: type switch (despacho por tipo concreto)
func PrintAny(x any) {
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case Abser:
		fmt.Println("Abser:", v.Abs())
	default:
		fmt.Printf("otro (%T): %v\n", v, v)
	}
}

func main() {
	// --- Métodos: valor vs puntero
	v := Vertex{3, 4}
	fmt.Println("Abs(valor):", v.Abs()) // 5
	v.Scale(2)                          // método puntero: Go usa &v implícito
	fmt.Println("Scale →", v)           // {6 8}
	ScalePtr(&v, 0.5)                   // función puntero: pasar &v explícito
	fmt.Println("ScalePtr →", v)        // {3 4}
	ScaleVal(v, 10)                     // no muta (copia)
	fmt.Println("ScaleVal (no muta) →", v)

	// --- Mutar con receiver puntero
	f := MyFloat(-2)
	fmt.Println("f start:", f)                 // -2
	fmt.Println("AbsInPlacePlus1:", (&f).AbsInPlacePlus1()) // 3 (muta f a |−2|+1)
	fmt.Println("f after:", f)                 // 3
	fmt.Println("AbsVal (no muta):", f.AbsVal())

	// --- Interfaces: implementación implícita + method set
	var a Abser
	mf := MyFloat2(-3)
	a = mf            // ok: MyFloat2 tiene Abs por valor
	fmt.Println("Abser(MyFloat2):", a.Abs())

	a = &v            // ok: *Vertex implementa Abser (Vertex no)
	fmt.Println("Abser(*Vertex):", a.Abs())

	// Interface con puntero nil dentro (a != nil)
	var pv *Vertex = nil
	a = pv
	fmt.Println("a == nil ? →", a == nil) // false (tipo dinámico presente)

	// --- Type assertions
	var i any = "hola"
	s, ok := i.(string) // segura
	fmt.Println("assert string ok?", ok, "→", s)

	// pánico controlado (no ejecutar en prod):
	// _ = i.(int) // panic: no es int

	// Type switch
	PrintAny(42)
	PrintAny("hi")
	PrintAny(&v)
}
