package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 6}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON) // Saída em bytes
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Mel",
		"raca": "SRD",
	}
	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

	cachorro3EmJSON := `{"nome":"Rex","raca":"Dalmata","idade":6}`
	var c3 cachorro
	if erro := json.Unmarshal([]byte(cachorro3EmJSON), &c3); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c3)

	cachorro4EmJSON := `{"nome":"Toby","raca":"Poodle"}`
	c4 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro4EmJSON), &c4); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c4)

}
