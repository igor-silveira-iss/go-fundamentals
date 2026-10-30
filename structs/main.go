package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

type deputado struct {
	nome    string
	partido string
	numero  uint16
}

func (d deputado) apresentacao() string {
	return d.nome + " (" + d.partido + ") "
}

func (d deputado) trocaPartido(novoPartido string) {
	d.partido = novoPartido
	fmt.Println("O novo partido é:", d.partido)
}

func main() {
	// fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	ende := endereco{"Rua dos Bobos", 66}

	u2 := usuario{"João", 45, ende}
	fmt.Println(u2)

	u3 := usuario{nome: "Cleiton"}
	fmt.Println(u3)

	// Teste "métodos"
	dep := deputado{"Lula", "PT", 78}
	fmt.Println(dep)
	apresentacaoDeputado := dep.apresentacao()
	fmt.Println(apresentacaoDeputado)
	dep.trocaPartido("MBL")

}
