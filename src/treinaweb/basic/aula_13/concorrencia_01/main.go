package main

import "fmt"

func main() {
	fmt.Println(" Concorrência - Programação Assíncrona")
	fmt.Println(" 1º Vamos de Programação síncrona")
	fmt.Println(" Ele vai contar o conteAte e depois o main ")

	var limite int
	fmt.Print("Informe um limite: ")
	fmt.Scanf("%d", &limite)
	conteAte(limite)
	for i := 0; i <= limite*10; i++ {
		fmt.Printf(" - |main| O número é %d \n", i)
	}
}

func conteAte(limite int) {
	for i := 0; i <= limite*20; i++ {
		fmt.Printf(" - |conteAte| O número é %d \n", i)
	}
}
