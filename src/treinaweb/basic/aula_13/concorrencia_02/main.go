package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" Concorrência - Programação Assíncrona")
	fmt.Println(" Ele vai contar o conteAte e  main ")

	var limite int
	fmt.Print("Informe um limite: ")
	fmt.Scanf("%d", &limite)
	/*
		conteAte agora recebe uma goroutine com a palavra chave go na frente
		Goroutines run in the same address space, so access to shared memory must be synchronized.
	*/
	go conteAte(limite)
	for i := 0; i <= limite*10; i++ {
		fmt.Printf(" - |main| O número é %d \n", i)
	}
	/*
		Multiplicando 10x a definição de 1 segundo.
		Objeto é conseguir visualizar o processamento Assíncrono
	*/
	time.Sleep(10 * time.Second)
}

// conteAte o limite
func conteAte(limite int) {
	for i := 0; i <= limite*20; i++ {
		fmt.Printf(" - |conteAte| O número é %d \n", i)
	}
}
