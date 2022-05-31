package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(200)
	fmt.Println(contaExemplo.ObterSaldo())

}
