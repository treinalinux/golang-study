package main

import (
	"alura/oo/Interface/contas"
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

	contaDoAlan := contas.ContaPoupanca{}
	contaDoAlan.Depositar(100)
	contaDoAlan.Sacar(15)

	fmt.Println("Saldo em conta do Alan:\t R$", contaDoAlan.ObterSaldo())

	contaDoCarla := contas.ContaCorrente{}
	contaDoCarla.Depositar(100)
	contaDoCarla.Sacar(10)

	fmt.Println("Saldo em conta Carla:\t R$", contaDoCarla.ObterSaldo())
}
