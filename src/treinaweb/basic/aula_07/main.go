package main

import "fmt"

func main() {
	fmt.Println("Funções e escopos:")
	var numero1, numero2 int
	var operacao string
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanf("%d", &numero1)
	fmt.Print("Digite a operação (+, -, *, /, $): ")
	fmt.Scanf("%s", &operacao)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanf("%d", &numero2)
	if operacao == "+" {
		somar(numero1, numero2)
	} else if operacao == "-" {
		resultado := subtrair(numero1, numero2)
		fmt.Printf("%d - %d = %d \n", numero1, numero2, resultado)
	} else if operacao == "*" {
		resultado := multiplicar(numero1, numero2)
		fmt.Printf("%d x %d = %d \n", numero1, numero2, resultado)
	} else if operacao == "/" {
		resultado, resto := dividir(numero1, numero2)
		fmt.Printf("QUOCIENTE = %d; RESTO = %d \n", resultado, resto)
	} else if operacao == "$" {
		incremento, decremento := incrementoDecremento(numero1, numero2)
		fmt.Printf("INCREMENTO = %d, DECCREMENTO = %d \n", incremento, decremento)
	} else {
		fmt.Println("Operação inválida!")
	}
}

// somar - uma função sem retorno
func somar(n1 int, n2 int) {
	fmt.Printf("%d + %d = %d \n", n1, n2, n1+n2)
}

// subtrair - uma função com retorno
func subtrair(n1 int, n2 int) int {
	return n1 - n2
}

// multiplicar - uma função com retorno sem preciar definir no corpo da função
func multiplicar(n1 int, n2 int) (resultado int) {
	resultado = n1 * n2
	return
}

// dividir - uma função com retorno de multiplos valores
func dividir(n1 int, n2 int) (int, int) {
	quociente := n1 / n2
	resto := n1 % n2
	return quociente, resto
}

// incrementoDecremento - uma função com retorno de multiplos valores
func incrementoDecremento(n1 int, n2 int) (inc int, dec int) {
	inc = n1 + n2
	dec = n1 - n2
	return
}
