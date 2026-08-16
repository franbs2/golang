package main

import "fmt"

func somar(a int8, b int8) int8 {
	return a + b
}

func calculos(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main(){
	soma := somar(1, 3)
	fmt.Println(soma)

	var f = func (txt string) string {
		fmt.Println(txt)
		return "resultado"
	}

	resultado :=f("Texto")
	fmt.Println(resultado)

	resultsoma, resultsub := calculos(15,8)

	fmt.Println(resultsoma, resultsub)

	resultsoma, _ = calculos(16,8)

	fmt.Println(resultsoma)
}