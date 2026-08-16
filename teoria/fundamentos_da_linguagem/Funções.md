Funções são um tipo um pouco diferente em Go.

`func` inicia a declaração de uma função.

As funções podem:

- receber parâmetros;
- retornar valores;
- receber vários parâmetros;
- retornar vários valores.

---

## Estrutura básica

```
func nome(nome_atr tipo) tipo {
    return nome_atr
}
```

Exemplo:

```
func dobrar(numero int) int {
    return numero * 2
}
```

Chamando a função:

```
resultado := dobrar(5)

fmt.Println(resultado)
```

---

# Funções anônimas

Também podemos armazenar uma função dentro de uma variável.

```
var f = func(txt string) {
    fmt.Println(txt)
}

f("Texto")
```

Nesse caso, `f` recebe uma função.

A função não possui um nome próprio. Por isso, é chamada de **função anônima**.

---

## Função anônima com retorno

Também podemos definir o tipo de retorno:

```
var f = func(txt string) string {
    fmt.Println(txt)
    return "resultado"
}

resultado := f("Texto")

fmt.Println(resultado)
```

---

# Múltiplos retornos

Uma função em Go pode retornar mais de um valor.

```
func calculos(n1, n2 int8) (int8, int8) {
    soma := n1 + n2
    sub := n1 - n2

    return soma, sub
}
```

Podemos receber os dois valores:

```
resultadoSoma, resultadoSub := calculos(15, 8)

fmt.Println(resultadoSoma, resultadoSub)
```

Também podemos ignorar um dos retornos usando `_`:

```
resultadoSoma, _ := calculos(16, 8)

fmt.Println(resultadoSoma)
```

O `_` é utilizado quando não queremos utilizar determinado valor retornado.

> A ordem dos retornos é importante. O primeiro valor retornado será atribuído à primeira variável, o segundo à segunda variável, e assim por diante.

---

# Funções variádicas

Uma função variádica aceita um número variável de argumentos.

Utilizamos `...` antes do tipo:

```
func somar(numeros ...int) int {
    soma := 0

    for _, numero := range numeros {
        soma += numero
    }

    return soma
}
```

Podemos passar quantos argumentos quisermos:

```
fmt.Println(somar(1, 2))
fmt.Println(somar(1, 2, 3, 4, 5))
```

Dentro da função, `numeros` funciona como um `slice`.

---

# Closures em Go

**Closures** são funções que conseguem capturar e manter acesso a variáveis do ambiente externo onde foram criadas.

Exemplo:

```
func contador() func() int {
    numero := 0

    return func() int {
        numero++
        return numero
    }
}
```

Podemos criar um contador:

```
contar := contador()

fmt.Println(contar())
fmt.Println(contar())
fmt.Println(contar())
```

Resultado:

```
1
2
3
```

A função anônima continua tendo acesso à variável `numero`, mesmo depois que a função `contador()` terminou sua execução.

Isso acontece porque a função retornada **capturou** a variável `numero`.

---

# Resumo

|Conceito|Significado|
|---|---|
|`func`|Declara uma função|
|Parâmetros|Valores recebidos pela função|
|Retorno|Valor devolvido pela função|
|Múltiplos retornos|Função pode retornar vários valores|
|`_`|Ignora um valor|
|`...`|Permite quantidade variável de argumentos|
|Função anônima|Função sem nome|
|Closure|Função que captura variáveis do ambiente externo|

---

# Material de leitura

- [Golang Functions — Medium](https://medium.com/@habbema/golang-functions-5cc9bbbb11f3)
- [Go: como criar e utilizar funções — Rocketseat](https://www.rocketseat.com.br/blog/artigos/post/go-como-criar-utilizar-funcoes)