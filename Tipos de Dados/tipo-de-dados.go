package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 100
	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	//alias
	//init32 = rune
	var numero3 rune = 1234
	fmt.Println(numero3)

	// byte = unit8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numero5 float32 = 123.4
	fmt.Println(numero5)

	var numero6 float64 = 123.43
	fmt.Println(numero6)

	numero7 := 13423.4
	fmt.Println(numero7)

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
