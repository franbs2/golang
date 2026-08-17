package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	endereco string
	numero   uint8
}

type pessoa struct {
	nome  string
	idade uint8
	peso  uint8
}

type estudante struct {
	pessoa
	altura uint8
	curso  string
}

func main() {
	fmt.Println("Arquivo struct")

	var usuario1 usuario = usuario{
		nome:  "Fran",
		idade: 20,
	}

	fmt.Println(usuario1)

	enderecousuario := endereco{"Rua dos bobos", 0}

	u2 := usuario{
		"David",
		20,
		enderecousuario,
	}

	var u usuario
	u.nome = "David"
	u.idade = 21

	fmt.Println(u)
	fmt.Println(u2)

	usuario3 := usuario{
		nome: "Davi",
	}

	fmt.Println(usuario3)

	p1 := pessoa{"joao", 20, 52}

	est := estudante{
		p1, 160, "eng",
	}

	fmt.Println(est)

}
