// === Tipos en Go (no hay clases) ===

// 1) STRUCT: tipo compuesto con campos
type User struct {
	ID   int
	Name string
}

// 2) Tipo DEFINIDO: nuevo identificador basado en otro tipo (NO es alias)
type Age int

// Método sobre tipo definido por el usuario
func (a Age) IsAdult() bool { return a >= 18 }

// 3) Alias de tipo: mismo tipo, otro nombre
type Rune = int32

// 4) Interface: contrato de métodos (implementación implícita; no hay "implements")
type Greeter interface{ Greet() string }

func (u User) Greet() string { return "Hi " + u.Name }

// === Type assertions (aseveraciones de tipo) ===

type Describer interface{ Describe() string }

type Person struct{ Name string }
func (p Person) Describe() string { return "Person: " + p.Name }

// Ejemplos de aserciones
func assertionsDemo() {
	var d Describer
	d = Person{"Valen"} // interface guarda (valor=Person{"Valen"}, tipo=Person)

	// Aserción "dura": si falla → panic
	_ = d.(Person)

	// Aserción "segura" (ok-idiom): no panic en fallo
	if p2, ok := d.(Person); ok {
		_ = p2.Name
	}

	// Falla controlada (tipo incorrecto) → ok=false, cero de T
	if _, ok := d.(string); !ok {
		// no panic
	}

	// Coincidencia exacta: si d guarda *Person, d.(Person) falla; usar d.(*Person)
}

// === Type switch (ramificar por tipo dinámico) ===

func PrintAny(x any) {
	switch v := x.(type) {
	case int:
		_ = v
	case string:
		_ = v
	case Describer:
		_ = v.Describe()
	default:
		_ = v // otro tipo
	}
}

// === Interface no nil con valor nil (caso especial) ===

func nilInterfaceNote() {
	var i Describer
	var pp *Person = nil
	i = pp           // (nil, *Person) → i != nil
	_ = (i == nil)   // false
	// Llamar i.Describe() puede panickear si el método no maneja receptor nil.
}

/*
Resumen clave:
- Go no tiene clases; trabajás con TIPOS (structs, tipos definidos) + métodos + interfaces.
- Implementación de interfaces es implícita: si el tipo tiene los métodos, ya implementa la interface.
- Una interface guarda (valor, tipo-concreto). Aserciones:
    t := i.(T)       // panic si no es T
    t, ok := i.(T)   // ok=false si no es T (sin panic)
- Type switch para ramificar por tipo.
- Coincidencia exacta: Person vs *Person.
*/
