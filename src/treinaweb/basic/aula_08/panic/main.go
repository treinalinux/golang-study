package main

import "fmt"

func main() {
	fmt.Println("Entendo o panic:")
	defer fmt.Println("Obrigado por usar nosso progrma.")
	fmt.Print("Digite um número maior que 10: ")
	var numero int
	fmt.Scanf("%d", &numero)
	if numero <= 10 {
		panic("Número inválido!")
	} else {
		fmt.Println("Você escolheu um bom número!")
	}

}
