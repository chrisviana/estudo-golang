package main

import "fmt"

func main() {
	var variave1 string = "Variavel 1"
	variavel2 := "Variavel2"
	fmt.Println(variave1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "leleleel"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "var 5", "var 6"
	fmt.Println(variavel5, variavel6)

	const cons1 string = "Const 1"
	fmt.Println(cons1)

}
