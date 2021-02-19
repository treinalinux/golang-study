package main

import "fmt"

func main() {
	fmt.Println("Ponteiros, exemplo sem ponteiro")
	var saldo float64
	fmt.Print("Digite o seu saldo: ")
	fmt.Scanf("%f", &saldo)
	saldo += calculaRendimento(saldo)
	fmt.Printf("Seu saldo com rendimentos é de R$ %.2f \n", saldo)
}

func calculaRendimento(saldo float64) float64 {
	// 3%
	return saldo * 0.03
}
