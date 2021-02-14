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

// Transferir entre contas no banco
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func main() {
	contaDoAlan := ContaCorrente{titular: "Alan", saldo: 300}
	contaDoCarla := ContaCorrente{titular: "Carla", saldo: 500}

	// Lembre do &, precisa de endereço para transferir
	status := contaDoCarla.Transferir(-200, &contaDoAlan)

	fmt.Println(status)

	fmt.Println(contaDoAlan)
	fmt.Println(contaDoCarla)
}
