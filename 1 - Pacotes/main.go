package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()

	// remover go mod tidy
	erro := checkmail.ValidateFormat("christgmail.com")
	fmt.Println(erro)
}
