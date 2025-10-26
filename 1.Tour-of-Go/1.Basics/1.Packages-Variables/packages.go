package main

//exista un paquete main con una función main()

//IMPORTANTE
//En Go no importa el nombre del archivo:
// lo importante es que el paquete sea main y exista func main()
//  para que go run . funcione.”
import (
	"fmt"
	"math/rand"
)

func main1() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

/*Si no hubiera func main() (o el paquete no fuera main),
fallaría con un error tipo: “go: cannot run non-main package” o
“function main is undeclared in the main package”.


Y si hubiera dos func main() en el mismo paquete, tendrías un “main redeclared”.*/
