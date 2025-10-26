# Hello World en Go (Debian 12 + VS Code)


## 1) Crear carpeta y módulo

```bash
# crea la carpeta del proyecto y entra
mkdir -p ./dev/hello && cd ./dev/hello

# inicializa el MÓDULO (Go usa módulos para gestionar deps y “saber quién sos”)
go mod init example.com/hello
```

> **¿Qué es `go.mod`?** Un archivo que define el nombre del módulo y sus dependencias.
> Si después subís el repo a GitHub, podés renombrarlo a `github.com/tuusuario/hello`.

---

## 2) Crear el archivo `main.go`

```bash
cat > main.go <<'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Hola, Go")
}
EOF
```

> **Claves:**
>
> * `package main` indica que es un ejecutable.
> * `func main()` es el punto de entrada del programa.
> * `fmt.Println` imprime texto en consola.

---

## 3) Ejecutar (rápido) vs Compilar (binario)

### Ejecutar “en caliente”

```bash
go run .
```

* **Qué hace:** compila en memoria y corre al toque. Ideal para probar.

### Compilar binario

```bash
mkdir -p bin
go build -o bin/hello .
./bin/hello
```

* **Qué hace:** genera un ejecutable en `bin/hello` y lo corrés directo.

---

## 4) (Opcional) Formatear y testear

### Formateo

```bash
gofmt -w .
```

> En VS Code, activá **Format on Save** y listo.

### Test mínimo (si querés ver cómo se prueba)

```bash
cat > main_test.go <<'EOF'
package main

import "testing"

func TestDummy(t *testing.T) {
    // test de ejemplo que siempre pasa
}
EOF

go test ./...
```

---

## 5) VS Code (extensión Go) en 30s

1. Abrí **VS Code** → instalá la extensión **“Go” (golang.go)**.
2. Aceptá instalar las herramientas sugeridas (gopls, gofmt/goimports, etc.).
3. Abrí la carpeta `~/dev/hello`, modificá `main.go`, guardá y probá `go run .`.

> **Depurar (opcional):** poné un breakpoint y presioná **F5**. Si te pide Delve, instalalo con
> `go install github.com/go-delve/delve/cmd/dlv@latest` y asegurate de tener `~/go/bin` en tu `PATH`.

---

## 6) Atajos mentales (resumen)

* `go mod init ...` → crear módulo.
* `go run .` → ejecutar rápido.
* `go build -o bin/hello .` → compilar binario.
* `gofmt -w .` → formatear código.
* `go test ./...` → correr tests (si hay).

---

## 7) Errores típicos

* **`go: no Go files in .`** → faltan `.go` o estás en la carpeta equivocada.
* **`command not found: go`** → faltó instalar Go o agregarlo al `PATH`.
* **No se formatea al guardar** → activá “Format on Save” y verificá la extensión Go.

---


