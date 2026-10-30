package main

import "fmt"

func main() {
	// Operadores Aritméticos
	soma := 1 + 2
	subtracao := 5 - 2
	divisao := 10 / 5
	multiplicacao := 2 * 9
	resto := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	// Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// Unários
	numero := 10
	numero++
	fmt.Println(numero)
	numero -= 3
	fmt.Println(numero)
}
