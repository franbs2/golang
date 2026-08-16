package main

import (
	"errors"
	"fmt"
)

func main() {
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "var 3"
		variavel4 string = "var 4"
	)
	fmt.Println(variavel3, variavel4)

	var (
		variavel5 string
		variavel6 string
	)
	variavel5, variavel6 = "var 5", "var 6"
	fmt.Println(variavel5, variavel6)

	variavel7, variavel8 := "var 7", "var 8"
	fmt.Println(variavel7, variavel8)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel1, variavel2 = variavel2, variavel1
	fmt.Println(variavel1, variavel2)

	// TIPOS DE DADOS

	var numero_inteiro int64 = 1000000000000
	fmt.Println(numero_inteiro)

	var numero_inteiro2 uint = 1
	println(numero_inteiro2)

	var numero3 rune = 123
	fmt.Println(numero3)

	var numero4 byte = 4
	fmt.Println(numero4)

	var numero_flutuante float32 = 12.5
	fmt.Println(numero_flutuante)

	var numero_flutuante2 float64 = 10000000000000000000002.5
	fmt.Println(numero_flutuante2)

	var str string = "uma string"
	fmt.Println(str)

	char := 'A'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)

	var erro_personalizado error = errors.New("Erro personalizado")
	fmt.Println(erro_personalizado)
}
