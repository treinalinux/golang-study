package main

import (
	"fmt"
)

func main() {
	// Strings
	fmt.Println("Strings:")
	var s1 string = "Alan da Silva Alves"
	fmt.Println("A variavel s1 contém o valor:", s1)
	var s2 string = `
	Alan da Silva Alves`
	fmt.Println("A variavel s2 contém o valor:", s2)
	// Código unicode do caracter
	fmt.Println("A variavel s2 contém o valor:", s2[0])
	fmt.Println("")

	// unit
	fmt.Println("Inteiros sem sinal:")
	var u1 uint8 = 255 // unit vai de 0-255
	fmt.Println(u1)
	var u2 uint16 = 33
	fmt.Println(u2)
	var u3 uint32 = 44
	fmt.Println(u3)
	var u4 uint64 = 55
	fmt.Println(u4)
	fmt.Println("")

	// int8
	fmt.Println("Inteiros com sinal:")
	var i1 int8 = 127
	fmt.Println(i1)
	var i2 int16 = 1000
	fmt.Println(i2)
	var i3 rune = 1001 // rune é um alias para int32
	fmt.Println(i3)
	var i4 int32 = 1002
	fmt.Println(i4)
	fmt.Println("")

	// uint depende da arquitetoria do processador
	var t1 uint = 1010
	fmt.Println(t1)
	var t2 int = 2000
	fmt.Println(t2)
	fmt.Println("")

	// float
	fmt.Println("Pontos flutuantes:")
	var f1 float32 = 1
	fmt.Println(f1)
	var f2 float64 = 2
	fmt.Println(f2)
	var f3 complex64 = 3
	fmt.Println(f3)
	var f4 complex128 = 4
	fmt.Println(f4)
	fmt.Println("")

	// Booleanos
	fmt.Println("Booleanos:")
	var b1 bool = true
	fmt.Println(b1)
	fmt.Println("")

	// Multiplas declarações
	fmt.Println("Multiplas declarações:")
	var nome, sobrenome string = "Alan", "da Silva Alves"
	fmt.Println(nome + " " + sobrenome)
	var (
		nome2      string = "Carla"
		sobrenome2 string = "Alves"
		idade      int    = 30
	)
	fmt.Println(nome2 + " " + sobrenome2)
	fmt.Println(idade, "anos")
	fmt.Println("")

	// Inferencia de tipos
	var nome3 = "Jessica"
	// Error para alguns tipos, como o tipo int
	// var idade2 = 30
	// Você pode omitir o tipo de váriavel com :=
	idade2 := 31
	fmt.Println(nome3, "sua idade é", idade2)
	fmt.Println("")

	// Constante
	const constante string = "Alan da Silva Alves"
	// cannot assign to constante (constant "Alan da Silva Alves" of type string)
	// constante = "Outro nome"
	fmt.Println(constante)

	const (
		primeiroNome = "Alan"
		ultimoNome   = "Alves"
	)

	fmt.Println(primeiroNome, ultimoNome)

}
