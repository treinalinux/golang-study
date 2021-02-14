package contas

import (
	"alura/oo/Interface/clientes"
)

// ContaCorrente foi criada para struct da conta corrente
type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
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

// ObterSaldo da conta Corrente
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
