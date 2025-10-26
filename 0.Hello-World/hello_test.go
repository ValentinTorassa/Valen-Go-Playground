package main

import "testing"

// t *testing.T es el manejador del test: reporta logs/errores, controla flujo (FailNow/Skip/Run), y mantiene el estado del test.
func TestHello(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"nombre normal", "valen", "Hola soy valen"},
		{"con mayúsculas", "Valen", "Hola soy Valen"},
		{"vacío", "", "Hola soy "},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Hello(tc.in)
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}
