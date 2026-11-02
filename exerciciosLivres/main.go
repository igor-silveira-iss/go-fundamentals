package main

import "fmt"

type cachorro struct {
	nome  string
	raca  string
	peso  float32
	idade uint16
}

func main() {
	// Atribuição de variáveis
	numero := 23
	fmt.Println(numero)
	texto := "teste"
	fmt.Println(texto)

	// Instanciando um struct
	cachorro1 := cachorro{
		"Valter",
		"Shitzu",
		5.5,
		3,
	}
	fmt.Println(cachorro1)

	// Criando um slice
	var nomes []string
	nomes = append(nomes, "Joaquim", "João Pedro", "Osvaldo", "Igor")
	fmt.Println(nomes)
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	idades := make([]uint16, 0, 12)
	idades = append(idades, 54, 67)
	fmt.Println(idades)

	// Criando um array
	arrayNomes := [5]string {"João", "Afonso", "Wilson", "Ulisses", "Humberto"}
	fmt.Println(arrayNomes)

	// Criando ponteiros
	var nomeCompleto string
	var ponteiroNome *string
	// fmt.Println(ponteiroNome)
	nomeCompleto = "Igor Silveira da Silva"
	// fmt.Println(nomeCompleto)
	ponteiroNome = &nomeCompleto
	fmt.Println(ponteiroNome, *ponteiroNome)
	nomeCompleto = "Sílvio Santos"
	fmt.Println(ponteiroNome, *ponteiroNome)

	// Criando Maps
	pessoas := map[string]string {
		"nome": "Ronaldo",
		"sobrenome": "Assis",
	}
	fmt.Println(pessoas)
	fmt.Println(pessoas["sobrenome"])

	// Loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
