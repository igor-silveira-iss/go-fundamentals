package main

import "fmt"

func main() {

	// var variavel1 int16 = 10
	// var variavel2 int16 = variavel1
	// fmt.Println(variavel1, variavel2)
	// variavel1++
	// fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int16
	var ponteiro *int16
	fmt.Println(variavel3, ponteiro)
	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro, *ponteiro)
	variavel3 = 150
	fmt.Println(variavel3, ponteiro, *ponteiro)

}
