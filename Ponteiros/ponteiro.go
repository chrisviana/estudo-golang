package main

import "fmt"

func main() {

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)
	v1++
	fmt.Println(v1, v2)

	//ponteiro é uma referencia de memoria
	var v3 int
	var ponteiro *int

	v3 = 100
	ponteiro = &v3

	fmt.Println(v3, ponteiro)

	v3 = 150
	fmt.Println(v3, *ponteiro) //desreferenciação

}
