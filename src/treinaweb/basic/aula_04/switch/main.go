package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estruturas condicionais com switch")

	var mes int
	fmt.Print("Digite o número do mês: ")
	fmt.Scanf("%d", &mes)
	switch mes {
	case 1:
		fmt.Println("Esse mês é Janeiro.")

	case 2:
		fmt.Println("Esse mês é Fevereiro.")

	case 3:
		fmt.Println("Esse mês é Março.")

	case 4:
		fmt.Println("Esse mês é Abril.")

	case 5:
		fmt.Println("Esse mês é Maio.")

	case 6:
		fmt.Println("Esse mês é Junho.")

	case 7:
		fmt.Println("Esse mês é Julho.")

	case 8:
		fmt.Println("Esse mês é Agosto.")

	case 9:
		fmt.Println("Esse mês é Setembro.")

	case 10:
		fmt.Println("Esse mês é Outubro.")

	case 11:
		fmt.Println("Esse mês é Novembro.")

	case 12:
		fmt.Println("Esse mês é Dezembro.")

	default:
		fmt.Println("Mês inválido!")
	}
	switch mes {
	case 1, 2, 3:
		fmt.Println("Primeiro Trimestre.")
	case 4, 5, 6:
		fmt.Println("Segundo Trimestre.")
	case 7, 8, 9:
		fmt.Println("Terceiro Trimestre.")
	case 10, 11, 12:
		fmt.Println("Quarto Trimestre.")
	}
}
