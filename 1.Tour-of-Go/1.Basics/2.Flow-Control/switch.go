package main

import (
	"fmt"
	"runtime"
	"time"
)

func main5() {
	fmt.Print("Go runs on ")

	// runtime.GOOS: constante de compilación con el SO objetivo
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	case today + 6:
		fmt.Println("In six days")
	default:
		fmt.Println("Too Far Away")
	}
	// time.Saturday es 6; comparamos la distancia (0..6) con módulo 7
	// para que “mañana” sea 1, etc.

	// Diferencia: un "switch <expr>" compara cada case con <expr> usando ==;
	// un "switch" sin expresión equivale a "switch true" y evalúa condiciones booleanas,
	// entrando en el **primer** case que resulte verdadero (útil para rangos).
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
