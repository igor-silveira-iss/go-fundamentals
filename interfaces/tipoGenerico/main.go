package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Texto")
	generica(10)
	generica(true)
	generica(23.87)
}
