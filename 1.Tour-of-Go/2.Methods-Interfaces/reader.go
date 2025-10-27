package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	src := strings.NewReader("Hola, io.Reader en Go!") // cualquier cosa que implemente Reader sirve igual
	buf := make([]byte, 8)                              // buffer de 8 bytes

	for {
		n, err := src.Read(buf) // intenta llenar buf; n puede ser < len(buf)
		if n > 0 {
			fmt.Printf("chunk: %q\n", buf[:n]) // procesás solo lo leído
		}
		if err == io.EOF { // fin del stream
			break
		}
		if err != nil { // error real (red, disco, etc.)
			fmt.Println("read error:", err)
			break
		}
	}
}
// Nota: Read puede devolver (n>0, err==io.EOF) en la última iteración; siempre procesá n antes de chequear el error.
