package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Alan"
	idade := 36
	version := 0.1
	fmt.Println("A variavel nome é do tipo:", reflect.TypeOf(nome))
	fmt.Println("Olá, sr.", nome+", a versão do programa é", version)
	fmt.Println("Olá, sr. tem", idade, "anos")
}
