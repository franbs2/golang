Go não é uma linguagem orientada a objetos tradicional e **não possui classes**. O recurso mais próximo de uma classe em Go é a **struct**.

Uma `struct` é uma coleção de campos com nome e tipo. Com ela, criamos um tipo de dado personalizado para agrupar informações relacionadas.

## Definindo uma struct

Para criar uma struct, definimos um tipo específico utilizando a palavra-chave `type`, o nome da estrutura e a palavra-chave `struct`:

```go
type nome struct {
    nome tipo
}
```

Exemplo:

```go
type usuario struct {
    nome  string
    idade uint8
}
```

---

## Exemplos de como usar

Existem diferentes formas de declarar, instanciar e atribuir valores a uma struct.

### 1. Declaração explícita com nomeação de campos

Informamos o tipo explicitamente e passamos os valores especificando o nome de cada campo:

```go
var usuario1 usuario = usuario{
    nome:  "Fran",
    idade: 20,
}

fmt.Println(usuario1)
```

### 2. Declaração e atribuição campo a campo

Declaramos a variável vazia do tipo da struct e preenchemos os valores utilizando o ponto (`.`):

```go
var u usuario
u.nome = "David"
u.idade = 21

fmt.Println(u)
```

### 3. Declaração com inferência pela ordem dos campos

Podemos passar os valores diretamente na ordem em que foram definidos na struct:

```go
u2 := usuario{
    "David",
    20,
}

fmt.Println(u2)
```

### 4. Inicialização parcial de campos

Quando especificamos os nomes dos campos, podemos preencher apenas os campos que quisermos. Os campos omitidos recebem automaticamente seu valor zero:

```go
usuario3 := usuario{
    nome: "Davi",
}

fmt.Println(usuario3)
```

---

## Structs dentro de structs (aninhadas)

Podemos chamar e utilizar structs como tipos de campos dentro de outras structs:

```go
type endereco struct {
    endereco string
    numero   uint8
}

type usuario struct {
    nome     string
    idade    uint8
    endereco endereco
}
```

Exemplo de uso:

```go
enderecousuario := endereco{"Rua dos bobos", 0}

u2 := usuario{
    "David",
    20,
    enderecousuario,
}

fmt.Println(u2)
```

---

## Quase uma herança (Composição / Embedding)

Go não possui herança tradicional. Para obter um comportamento similar, usamos **composição** por meio de campos anônimos (*struct embedding*).

Ao colocar uma struct dentro de outra sem definir um nome de campo, todos os campos da struct embutida passam a fazer parte da nova struct:

```go
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
```

### Inicializando e acessando

Precisa passar a instância de `pessoa` para `estudante`. Na leitura, **não precisa chamar o nome da struct interna (`pessoa`)**:

```go
p1 := pessoa{"joao", 20, 52}

est := estudante{
    p1,
    160,
    "eng",
}

fmt.Println(est)

// Acesso direto aos campos de pessoa através de estudante:
fmt.Println(est.nome)
fmt.Println(est.idade)
fmt.Println(est.curso)
```

---

## Resumo

| Conceito | Descrição | Exemplo |
| :--- | :--- | :--- |
| **Definição** | Criação de um novo tipo estruturado | `type usuario struct { ... }` |
| **Atribuição nomeada** | Preenche apenas os campos passados | `usuario{nome: "Fran"}` |
| **Atribuição ordenada** | Preenche todos os campos em ordem | `usuario{"David", 20}` |
| **Struct aninhada** | Struct como campo de outra struct | `endereco endereco` |
| **Composição (Embedding)** | Embutir struct sem nome de campo | `type estudante struct { pessoa }` |

---

## Material para leitura

- [ttemporin.dev — O que são structs e como usá-las](https://ttemporin.dev/o-que-sao-structs-e-como-usa-las/)
- [DigitalOcean — Defining Structs in Go](https://www.digitalocean.com/community/tutorials/defining-structs-in-go)