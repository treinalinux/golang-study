package main

import (
	"fmt"
)

func main() {
	fmt.Println(" Concorrência - Programação Assíncrona")
	// Por padrão goroutines não pode compartilhar informação, por isso usuremos o channel
	fmt.Println(" Retornando valores com channel")

	var limite int
	canal1 := make(chan int)
	canal2 := make(chan int)
	fmt.Print("Informe um limite: ")
	fmt.Scanf("%d", &limite)
	/*
		conteAte agora recebe uma goroutine com a palavra chave go na frente
		Goroutines run in the same address space, so access to shared memory must be synchronized.
	*/
	go conteAte(limite, canal1)
	// função anônima
	go func(n int) {
		resultado := 0
		for i := 0; i <= n*10; i++ {
			fmt.Printf(" - |anônimo| O número é %d \n", i)
			resultado = i * 10
		}
		canal2 <- resultado
	}(limite)
	for i := 0; i <= limite*10; i++ {
		fmt.Printf(" - |main| O número é %d \n", i)
		/*
			Multiplicando 10x a definição de 1 segundo.
			Objeto é conseguir visualizar o processamento Assíncrono
		*/
	}
	// time.Sleep(10 * time.Second) # Uso não é maus necessário por causa do canal
	// resultadoCanal1 := <-canal1
	// resultadoCanal2 := <-canal2
	// fmt.Printf("Canal 1 = %d, Canal 2 %d \n", resultadoCanal1, resultadoCanal2)
	fmt.Printf("Canal 1 = %d, Canal 2 %d \n", <-canal1, <-canal2)
}

// conteAte o limite, agora com um canal de parametro
func conteAte(limite int, canal chan int) {
	resultado := 0
	for i := 0; i <= limite*20; i++ {
		fmt.Printf(" - |conteAte| O número é %d \n", i)
		resultado = i * 20
	}
	canal <- resultado
}
