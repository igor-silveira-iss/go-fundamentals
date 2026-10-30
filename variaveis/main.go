package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" // Declaração explícita
	fmt.Println(variavel1)

	variavel2 := "Variável 2" // Declaração implícita
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lelele"
	)
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5)
	fmt.Println(variavel6)

}
