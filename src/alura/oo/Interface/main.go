package main

import (
	"alura/oo/Interface/contas"
	"fmt"
	"os"
	"os/exec"
)

// PagarBoleto com verificar conta
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

// Interface criada
type verificarConta interface {
	Sacar(valor float64) string
}

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

	contaDoAlan := contas.ContaCorrente{}
	contaDoAlan.Depositar(100)
	PagarBoleto(&contaDoAlan, 30)
	fmt.Println("Alan, boleto pago, saldo atual:\t\t R$", contaDoAlan.ObterSaldo())

	contaDoCarla := contas.ContaPoupanca{}
	contaDoCarla.Depositar(100)
	PagarBoleto(&contaDoCarla, 90)
	fmt.Println("Carla, boleto pago, saldo atual:\t R$", contaDoCarla.ObterSaldo())

}
