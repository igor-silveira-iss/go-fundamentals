package main

import "fmt"

func inverterSinal(numero int32) int32 {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInveritdo := inverterSinal(int32(numero))
	fmt.Println(numeroInveritdo)

	novoNumero := 40
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
