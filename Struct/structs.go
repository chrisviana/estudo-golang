package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Struct")

	var u usuario
	u.nome = "Christian"
	u.idade = 24
	fmt.Println(u)

	endereco1 := endereco{"rua dos bobos", 0}
	u2 := usuario{"Christian", 24, endereco1}
	fmt.Println(u2)

	u3 := usuario{nome: "Christian"}
	fmt.Println(u3)
}
