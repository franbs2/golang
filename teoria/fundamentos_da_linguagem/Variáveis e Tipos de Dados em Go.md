
Todas as variáveis no Go possuem um **tipo**, seja ele declarado explicitamente ou inferido implicitamente.

---

## Declaração explícita

Quando informamos o tipo da variável:

```
var variavel tipo = valor
```

Exemplo:

```
var idade int = 20
```

---

## Declaração implícita

Podemos deixar o Go descobrir o tipo a partir do valor:

```
variavel := valor
```

Exemplo:

```
idade := 20
```

O tipo é determinado pelo valor atribuído.

> `:=` só pode ser usado dentro de funções.

---

# Declarando várias variáveis

Podemos declarar várias variáveis utilizando `var`:

```
var (
    nome string = "Fran"
    idade int = 20
)
```

Também podemos declarar sem atribuir valores:

```
var (
    nome string
    idade int
)
```

Também podemos atribuir vários valores de uma vez:

```
nome, idade = "Fran", 20
```

Ou declarar e atribuir utilizando `:=`:

```
nome, idade := "Fran", 20
```

---

# Constantes

Constantes são valores que não podem ser alterados depois de definidos.

```
const constante tipo = valor
```

Exemplo:

```
const pi float64 = 3.14159
```

Também podemos deixar o Go inferir o tipo:

```
const pi = 3.14159
```

---

# Invertendo valores de variáveis

Em Go, podemos inverter os valores de duas variáveis diretamente:

```
variavel1, variavel2 = variavel2, variavel1
```

Exemplo:

```
a := 10
b := 20

a, b = b, a
```

Resultado:

```
a = 20
b = 10
```

Não é necessário criar uma terceira variável para fazer a troca.

---

# Tipos de dados

Os principais tipos de dados em Go incluem:

- Números inteiros
- Números de ponto flutuante
- Números complexos
- Strings
- Booleanos
- Erros

---

# Números inteiros

Os tipos inteiros possuem diferentes tamanhos:

```
int8
int16
int32
int64
int
```

O número indica quantos **bits** são utilizados para representar o valor.

### `int`

O `int` possui tamanho dependente da arquitetura do computador.

Em geral:

- arquitetura 32 bits → `int` possui 32 bits;
    
- arquitetura 64 bits → `int` possui 64 bits.
    

Quando declaramos um número inteiro sem especificar o tipo, o Go normalmente utiliza `int`.

```
idade := 20
```

Nesse caso, `idade` é do tipo `int`.

---

# Inteiros sem sinal

Os tipos `uint` representam inteiros **sem sinal**, ou seja, não aceitam valores negativos.

Também existem:

```
uint8
uint16
uint32
uint64
uint
```

Exemplo:

```
var quantidade uint = 10
```

Não podemos fazer:

```
var quantidade uint = -10
```

---

# Aliases

Alguns tipos possuem aliases na linguagem Go:

```
int32  → rune
uint8  → byte
```

Por exemplo:

```
var letra rune = 'A'
```

É equivalente a:

```
var letra int32 = 'A'
```

---

# Números reais

Para representar números com casas decimais, utilizamos:

```
float32
float64
```

### `float32`

Utiliza 32 bits.

### `float64`

Utiliza 64 bits e possui maior precisão.

Exemplo:

```
var altura float64 = 1.80
```

Diferentemente de `int`, não existe um tipo `float` que escolha automaticamente entre `float32` e `float64` de acordo com a arquitetura.

Quando escrevemos:

```
altura := 1.80
```

o Go infere `float64`.

---

# Números complexos

Go também possui suporte a números complexos:

```
complex64
complex128
```

Exemplo:

```
var numero complex64
```

---

# Strings

`string` é utilizado para representar uma sequência de caracteres.

```
var nome string = "Fran"
```

Strings normalmente são escritas utilizando **aspas duplas**:

```
"Olá, mundo!"
```

Também podemos utilizar a inferência:

```
nome := "Fran"
```

---

# Caracteres

Go não possui um tipo `char` como algumas outras linguagens.

Um caractere entre aspas simples:

```
'A'
```

é representado por um valor inteiro e, normalmente, utilizamos o tipo `rune`.

```
var letra rune = 'A'
```

Como `rune` é um alias de `int32`, podemos obter o valor numérico correspondente ao caractere.

```
fmt.Println('A')
```

---

# Literais de string bruta

Strings entre **crases** são chamadas de _raw string literals_.

O conteúdo é interpretado praticamente da mesma forma como foi escrito, preservando quebras de linha e evitando a interpretação de sequências de escape.

```
texto := `Olá
mundo!`
```

É útil quando precisamos escrever textos com várias linhas ou que contenham muitos caracteres especiais.

---

# Literais de string interpretadas

Strings entre **aspas duplas** são chamadas de _interpreted string literals_.

Nelas, sequências especiais são interpretadas.

```
texto := "Olá\nmundo!"
```

O `\n` representa uma quebra de linha.

---

# Valores zero

Quando uma variável é declarada sem um valor inicial, o Go atribui automaticamente um **valor zero**, que depende do tipo.

Exemplos:

|Tipo|Valor zero|
|---|---|
|`int`|`0`|
|`float64`|`0`|
|`string`|`""`|
|`bool`|`false`|
|ponteiros|`nil`|
|slices|`nil`|
|maps|`nil`|
|interfaces|`nil`|

Exemplo:

```
var idade int
var nome string
var ativo bool
```

Os valores serão:

```
idade → 0
nome → ""
ativo → false
```

---

# Booleanos

O tipo `bool` possui apenas dois valores:

```
true
false
```

Exemplo:

```
var maiorDeIdade bool = true
```

Se uma variável `bool` não for inicializada, seu valor zero será:

```
false
```

---

# Error

`error` é uma interface utilizada pelo Go para representar erros.

É muito comum uma função retornar um valor e um `error`:

```
resultado, err := algumaFuncao()
```

Podemos verificar se ocorreu um erro:

```
if err != nil {
    fmt.Println("Ocorreu um erro:", err)
}
```

Em Go, erros são tratados explicitamente pelo programa, em vez de normalmente serem lançados como exceções.

---

# Resumo dos tipos

|   |   |
|---|---|
|Tipo|Uso|
|`int`|Números inteiros|
|`uint`|Inteiros sem sinal|
|`float32`|Números decimais|
|`float64`|Números decimais com maior precisão|
|`complex64`|Números complexos|
|`complex128`|Números complexos com maior precisão|
|`string`|Textos|
|`rune`|Caracteres / valores Unicode|
|`byte`|Alias de `uint8`|
|`bool`|`true` ou `false`|
|`error`|Representação de erros|
|`nil`|Ausência de valor em tipos que permitem `nil`|

---

# Material de leitura

- [W3Schools — Go Data Types](https://www-w3schools-com.translate.goog/go/go_data_types.php?_x_tr_sl=en&_x_tr_tl=pt&_x_tr_hl=pt&_x_tr_pto=tc)
- [DigitalOcean — Understanding Data Types in Go](https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go-pt)