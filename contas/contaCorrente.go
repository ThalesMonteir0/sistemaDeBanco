package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}

}
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "deposito realizado com sucesso. O saldo atual é:", c.saldo
	} else {
		return "valor menor que zero. O seu saldo é:", c.saldo
	}
}
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.saldo > valorDaTransferencia && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.saldo += valorDaTransferencia
		return true
	} else {
		return false
	}

}
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo

}
