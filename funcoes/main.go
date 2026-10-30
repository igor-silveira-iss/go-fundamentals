package main

import "fmt"

func somar(num1 int8, num2 int8) int8 {
	return num1 + num2
}

func calculosMatematicos(num1 int16, num2 int16) (int16, int16) {
	soma := num1 + num2
	subtracao := num1 - num2
	return soma, subtracao
}

func main() {
	var soma int8 = somar(10, 8)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadosSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadosSoma)
}
