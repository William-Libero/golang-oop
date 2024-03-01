package contas

import "../cliente"

type ContaCorrente struct {
	Titular       cliente.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := c.Saldo >= valor && valor > 0
	if podeSacar {
		c.Saldo -= valor
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	podeDepositar := valor > 0
	if podeDepositar {
		c.Saldo += valor
		return "Depósito realizado com sucesso! Saldo da conta:", valor
	} else {
		return "Valor de depósito insuficiente Saldo da conta:", valor
	}
}

func (contaRemetente *ContaCorrente) Transferencia(valor float64, contaDestino *ContaCorrente) string {
	podeTransferir := contaRemetente.Saldo >= valor && valor > 0
	if podeTransferir {
		contaRemetente.Saldo -= valor
		contaDestino.Depositar(valor)
		return "Transferencia realizado com sucesso!"
	} else {
		return "Valor de transferencia insuficiente."
	}
}
