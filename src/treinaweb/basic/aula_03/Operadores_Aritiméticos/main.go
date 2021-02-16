package main

import (
	"fmt"
)

func main() {
	fmt.Println("Operadores Aritiméticos")
	fmt.Println("")

	var numero1 int
	var numero2 int

	fmt.Print("Digite o primeiro número: ")
	// Posição de memória da variavel numero1 usando o &
	fmt.Scanln(&numero1)

	fmt.Print("Digite o segundo número:  ")
	// Posição de memória da variavel numero2 usando o &
	fmt.Scanln(&numero2)
	fmt.Println("")
	fmt.Printf("Soma..........: %d +  %d  = %d  \n", numero1, numero2, numero1+numero2)
	fmt.Printf("Subtração ....: %d -  %d  = %d  \n", numero1, numero2, numero1-numero2)
	fmt.Printf("Multiplicação.: %d *  %d  = %d  \n", numero1, numero2, numero1*numero2)
	fmt.Printf("Divisão.......: %d /  %d  = %d  \n", numero1, numero2, numero1/numero2)
	fmt.Printf("Resto.........: %d %%  %d  = %d  \n", numero1, numero2, numero1%numero2)
	fmt.Println("")

	incremento := numero1
	incremento += numero2
	fmt.Printf("Imcremento....: %d += %d é %d \n", numero1, numero2, incremento)

	decremento := numero1
	decremento -= numero2
	fmt.Printf("Decremento....: %d -= %d é %d \n", numero1, numero2, decremento)
}
