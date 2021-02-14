package main

import "fmt"

// ContaCorrente foi criada para struct da conta corrente
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// Sacar c de conta que estiver chamado o saque
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente!"
}

// Depositar criado para depositar na conta
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso! R$", c.saldo
	}
	return "Deposito falhou, pois é menor do que zero! R$", c.saldo
}

func main() {
	contaDoAlan := ContaCorrente{}
	contaDoAlan.titular = "Alan"
	contaDoAlan.saldo = 500

	fmt.Println(contaDoAlan.saldo)
	status, valor := contaDoAlan.Depositar(1600)
	fmt.Println(status, valor)

}
