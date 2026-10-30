package main

import (
	"errors"
	"fmt"
)

func main() {
	// int, int8, int16, int32, int64
	var numeroInteiro int16 = 100
	fmt.Println(numeroInteiro)

	// uint -> untyped int
	var numeroUint uint = 32
	fmt.Println(numeroUint)

	// alias
	var numeroAlias rune = 123456
	fmt.Println(numeroAlias)
	var numeroAlias2 byte = 123
	fmt.Println(numeroAlias2)

	// float32, float64
	var numeroPontoFlutuante float32 = 3.14
	fmt.Println(numeroPontoFlutuante)

	// string
	var nome string = "João da Silva"
	fmt.Println(nome)

	// Valor Zero
	var texto string
	fmt.Println(texto)

	// boolean
	var verdade bool = true
	fmt.Println(verdade)
	var falso bool = false
	fmt.Println(falso)

	// erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
