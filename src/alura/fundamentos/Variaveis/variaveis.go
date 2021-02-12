package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Alan"
	var idade int = 36
	var version float32 = 0.1
	fmt.Println("Olá, sr.", nome+", a versão do programa é", version)
	fmt.Println("Olá, sr. tem", idade, "anos")

	fmt.Println("O tipo da variavel é", reflect.Type0f(version))
}
