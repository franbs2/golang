package main

import (
	"fmt"
	"meu_projeto/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("franbs.ia@gmail.com")
	fmt.Println(erro)
}
