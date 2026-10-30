package main

import "fmt"

func main() {
	fmt.Println("Aula de Maps")
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"], usuario["sobrenome"])
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["ultimo"])
	delete(usuario2, "nome")
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}
