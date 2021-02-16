package main

import (
	"fmt"
	"strconv"
)

func main() {
	var texto string
	// Recebendo dados dos usuários
	fmt.Print("Digite um número: ")
	fmt.Scanf("%s", &texto)

	var numero1 int
	// Usamos o _ para ignorar o valor de erro que é retornado por Atoi()
	// O strconv é usado para converão de tipos
	numero1, _ = strconv.Atoi(texto)
	// Obs.: no new variables on left side of :=
	fmt.Println("O texto foi convertido para número 1:", numero1)

	// Outra maneira, vamos usar o decimal e 32 bits
	numero2, _ := strconv.ParseInt(texto, 10, 32)
	fmt.Println("O texto foi convertido para número 2:", numero2)

	// Outra maneira, é de forma direta, vamos usar o float64
	numero3, _ := strconv.ParseInt(texto, 10, 32)
	var convertidoParaFloat = float64(numero3)
	fmt.Println("O texto foi convertido para número 3:", convertidoParaFloat)

	fmt.Println(11 / 2)

}
