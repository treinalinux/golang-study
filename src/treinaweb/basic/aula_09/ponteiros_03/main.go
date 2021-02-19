package main

import "fmt"

func main() {
	fmt.Println("Ponteiros, exemplo com ponteiro")
	// var saldo float64
	// Declarando o ponteiro de for explicita usando o new(), já com a posição de memória
	var saldo = new(float64)
	fmt.Print("Digite o seu saldo: ")
	// %f de float
	// fmt.Scanf("%f", &saldo)
	fmt.Scanf("%f", saldo)
	// Fazendo um passagem por referência
	// calculaRendimento(&saldo)
	calculaRendimento(saldo)
	// %.2f para float
	// fmt.Printf("Seu saldo com rendimentos é de R$ %.2f \n", saldo)
	fmt.Printf("Seu saldo com rendimentos é de R$ %.2f \n", *saldo)
}

// calculaRendimento com uso do *ponteiro
func calculaRendimento(saldo *float64) {
	// 3%
	*saldo += *saldo * 0.03
}
