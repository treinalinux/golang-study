package main

import (
	"alura/oo/Composicao/contas"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//	contaDoAlan := contas.ContaCorrente{Titular: clientes.Titular{
	//		Nome:      "Alan",
	//		CPF:       "123.111.127.12",
	//		Profissao: "Desenvolvedor"},
	//		NumeroAgencia: 123, NumeroConta: 12234, Saldo: 1000}
	//
	//	fmt.Println(contaDoAlan)

	// Limpar a tela do terminal
	limpar := exec.Command("clear")
	limpar.Stdout = os.Stdout
	limpar.Run()

	fmt.Println("")

	contaDoAlan := contas.ContaCorrente{}
	contaDoAlan.Depositar(100)

	fmt.Println("Saldo em conta: R$", contaDoAlan.ObterSaldo())
}
