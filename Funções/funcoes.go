package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func calculos(a, b int) (int, int) {
	soma := a + b
	subtracao := a - b
	return soma, subtracao
}
func main() {
	soma := somar(1, 2)
	fmt.Println(soma)

	var resultado1, _ = calculos(1, 2)
	fmt.Println(resultado1)
}
