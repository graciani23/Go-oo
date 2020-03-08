package main

import (
	"fmt"

	"github.com/alura-Go/Banco/clientes"
	"github.com/alura-Go/Banco/contas"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.111.123.12", "Desenvolvedor"}
	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	fmt.Println(contaDoBruno)
}
