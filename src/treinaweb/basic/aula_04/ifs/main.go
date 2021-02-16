package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estruturas condicionais com if")

	var idade int
	fmt.Print("Digite a usa idade: ")
	fmt.Scanf("%d", &idade)
	// Outra maneira mais simples seria usar os dois && no primeiro if:
	// if idade >= 18 && idade <= 70
	// assim você pode omitir o else if
	if idade >= 18 && idade <= 70 {
		fmt.Println("Você pode dirigir!")
	} else {
		fmt.Println("Você não pode dirigir!")
	}
}
