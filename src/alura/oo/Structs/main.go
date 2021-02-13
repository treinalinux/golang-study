package main

import (
	"fmt"
)

// ContaCorrente foi criada para struct da conta corrente
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoAlan := ContaCorrente{titular: "Alan", numeroAgencia: 001, numeroConta: 789321, saldo: 50.5}

	contaDaCarla := ContaCorrente{"Carla", 001, 124987, 250.1}
	fmt.Println(contaDoAlan, contaDaCarla)

	// Aqui estamos usando ponteiros no go
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500.50

	// Use o * para retirar o & (local na memória) do resultado: &{Cris 0 0 500.5}
	fmt.Println(*contaDaCris)
	// Se tiver 2 contas da cris neste modelo o local na memória (&) será diferente mesmo
	// com conteúdos (* ponteiros) iguais.
	fmt.Println(&contaDaCris)
}
