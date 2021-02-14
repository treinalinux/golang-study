package main

import (
	"alura/oo/Retornos/Pacotes/contas"
	"fmt"
)

func main() {
	contaDoAlan := contas.ContaCorrente{Titular: "Alan", Saldo: 300}
	contaDoCarla := contas.ContaCorrente{Titular: "Carla", Saldo: 500}

	// Lembre do &, precisa de endereço para transferir
	status := contaDoCarla.Transferir(-200, &contaDoAlan)

	fmt.Println(status)

	fmt.Println(contaDoAlan)
	fmt.Println(contaDoCarla)
}
