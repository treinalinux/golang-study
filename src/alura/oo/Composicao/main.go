package main

import (
	"alura/oo/Composicao/clientes"
	"alura/oo/Composicao/contas"
	"fmt"
)

func main() {
	//	contaDoAlan := contas.ContaCorrente{Titular: clientes.Titular{
	//		Nome:      "Alan",
	//		CPF:       "123.111.127.12",
	//		Profissao: "Desenvolvedor"},
	//		NumeroAgencia: 123, NumeroConta: 12234, Saldo: 1000}
	//
	//	fmt.Println(contaDoAlan)
	clienteAlan := clientes.Titular{"Alan", "123.111.127.12", "Desenvolvedor"}
	contaDoAlan := contas.ContaCorrente{clienteAlan, 123, 12234, 1000}

	fmt.Println(contaDoAlan)
}
