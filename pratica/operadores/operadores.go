package main

import "fmt"

func main() {

	soma := 1 + 1
	sub := 2 - 1
	div := 10 / 4
	mul := 10 * 5
	resto := 10 % 2

	fmt.Println(soma, sub, div, mul, resto)

	var numero1 int16 = 10
	var numero2 int16 = 25

	soma2 := numero1 + numero2
	fmt.Println(soma2)

	var variavel1 string = "String"
	variavel2 := "string 2"

	fmt.Println(variavel1, variavel2)

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	numero10 := 10
	numero10++
	fmt.Println(numero10)

	numero10 += 15
	fmt.Println(numero10)

	numero10--
	numero10 -= 10

	numero10 *= 3
	numero10 %= 1

	fmt.Println(numero10)

}
