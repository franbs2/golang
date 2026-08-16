Operadores são símbolos que realizam operações matemáticas, lógicas, comparações e atribuições de valores entre variáveis e constantes.

## Operadores aritméticos

Utilizados para realizar cálculos matemáticos básicos.

Os operadores aritméticos disponíveis são:

`+` (adição ou concatenação de strings)  
`-` (subtração)  
`*` (multiplicação)  
`/` (divisão)  
`%` (módulo ou resto da divisão)  

Exemplo:

```go
soma := 10 + 5        // 15
subtracao := 10 - 5   // 5
multiplicacao := 10 * 5 // 50
divisao := 10 / 2     // 5
resto := 10 % 3       // 1
```

O operador `+` também pode ser usado para juntar textos:

```go
nomeCompleto := "Fran" + " " + "Silva"
```

### Operações entre tipos diferentes

Em Go, dois valores de tipos diferentes não podem ser operados diretamente.

O código abaixo não compila:

```go
var a int = 10
var b float64 = 2.5

resultado := a + b // Erro de compilação!
```

Para realizar o cálculo, é necessário converter explicitamente um dos tipos:

```go
resultado := float64(a) + b
```

## Operadores de atribuição

Utilizados para atribuir ou atualizar valores em variáveis.

### Atribuição simples

Utiliza o sinal `=`:

```go
variavel = valor
```

Exemplo:

```go
var idade int
idade = 20
```

### Atribuição com inferência de tipo (declaração curta)

Utiliza o operador `:=`:

```go
variavel := valor
```

Exemplo:

```go
idade := 20
```

`:=` declara a variável e já define seu tipo a partir do valor informado.  
Lembrando que `:=` só pode ser utilizado dentro de funções.

### Atribuição composta

Combinam uma operação aritmética com a atribuição de valor:

`+=` (soma e atribui)  
`-=` (subtrai e atribui)  
`*=` (multiplica e atribui)  
`/=` (divide e atribui)  
`%=` (calcula o resto e atribui)  

Exemplo:

```go
x := 10

x += 5 // x = x + 5 (resultado: 15)
x -= 3 // x = x - 3 (resultado: 12)
x *= 2 // x = x * 2 (resultado: 24)
x /= 4 // x = x / 4 (resultado: 6)
x %= 2 // x = x % 2 (resultado: 0)
```

## Operadores relacionais

Utilizados para comparar dois valores. O resultado de uma comparação é sempre um valor booleano (`true` ou `false`).

Os operadores relacionais são:

`==` (igual a)  
`!=` (diferente de)  
`>` (maior que)  
`<` (menor que)  
`>=` (maior ou igual a)  
`<=` (menor ou igual a)  

Exemplo:

```go
a := 10
b := 20

fmt.Println(a == b) // false
fmt.Println(a != b) // true
fmt.Println(a > b)  // false
fmt.Println(a < b)  // true
fmt.Println(a >= 10) // true
fmt.Println(b <= 20) // true
```

Assim como nas operações aritméticas, só podemos comparar valores do mesmo tipo.

## Operadores lógicos

Utilizados para combinar expressões booleanas.

### E lógico (AND)

Representado por `&&`.  
Retorna `true` apenas se todas as condições forem verdadeiras.

```go
idade := 20
temCarteira := true

podeDirigir := idade >= 18 && temCarteira
// true
```

### OU lógico (OR)

Representado por `||`.  
Retorna `true` se pelo menos uma das condições for verdadeira.

```go
temCupom := false
eAniversariante := true

temDesconto := temCupom || eAniversariante
// true
```

### Negação lógica (NOT)

Representado por `!`.  
Inverte o valor booleano: se for `true` vira `false`, e se for `false` vira `true`.

```go
ativo := true
bloqueado := !ativo
// false
```

## Operadores unários (incremento e decremento)

Utilizados para aumentar ou diminuir o valor de uma variável em 1.

`++` (incremento)  
`--` (decremento)  

Exemplo:

```go
contador := 0

contador++ // contador agora vale 1
contador-- // contador volta a valer 0
```

### Peculiaridades em Go

Diferente de outras linguagens como C, Java ou JavaScript:

Em Go, `++` e `--` só existem na forma pós-fixada. Não existe pré-fixado:

```go
++contador // Erro de sintaxe!
```

`++` e `--` são instruções (statements) e não expressões. Por isso, não podem ser atribuídos a outra variável:

```go
y := x++ // Erro de sintaxe!
```

## Resumo dos operadores

| Operador | Categoria | Descrição |
| :--- | :--- | :--- |
| `+`, `-`, `*`, `/`, `%` | Aritméticos | Cálculos matemáticos |
| `=` | Atribuição | Atribuição simples |
| `:=` | Atribuição | Declaração curta com inferência de tipo |
| `+=`, `-=`, `*=`, `/=`, `%=` | Atribuição composta | Operação aritmética e atribuição direta |
| `==`, `!=` | Relacionais | Igualdade e diferença |
| `>`, `<`, `>=`, `<=` | Relacionais | Comparações de magnitude |
| `&&` | Lógico | E (AND) |
| `\|\|` | Lógico | OU (OR) |
| `!` | Lógico | Negação (NOT) |
| `++`, `--` | Unários | Incremento e decremento |

## Material de leitura

- [Micilini — Operadores em Go](https://micilini.com/conteudos/golang/operadores-em-go)
- [Go by Example — Operators](https://gobyexample.com/)
- [W3Schools — Go Operators](https://www.w3schools.com/go/go_operators.php)