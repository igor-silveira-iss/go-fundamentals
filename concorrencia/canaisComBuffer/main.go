package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go"
	// canal <- "Simulando Deadlock"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
