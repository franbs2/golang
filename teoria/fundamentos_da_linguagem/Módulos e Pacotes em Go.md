Quando lidamos com mais de um pacote no Go, precisamos usar **módulos**. Os módulos são conjuntos de pacotes relacionados.

Usamos:

```go
go mod init [nome_do_modulo]
```

Esse comando cria o arquivo `go.mod`, que centraliza as dependências do projeto. Se houver alguma biblioteca externa, ela ficará registrada aqui, com seu nome e versão.

Depois de adicionarmos o código em `[praticas/pacotes/main.go]`, podemos rodar com:

```go
go run main.go
```

Ou usar:

```go
go build
```

O `go build` irá criar um arquivo executável, normalmente com o nome do módulo, que você definiu no `go mod init`.

Esse é o arquivo compilado do projeto. Podemos executar o projeto rodando o arquivo:

```bash
./[nome_do_modulo]
```

---

**Pacotes** são conjuntos de arquivos `.go` que ficam em uma mesma pasta e pertencem ao mesmo pacote.

Funções com a letra inicial **maiúscula** podem ser usadas em outros pacotes do projeto. Elas são como funções "públicas", ou seja, podem ser exportadas.

```go
func Somar() {
    // ...
}
```

Já funções com a letra inicial **minúscula** não podem ser acessadas diretamente por outros pacotes.

```go
func somar() {
    // ...
}
```

Dentro do mesmo pacote, podemos chamar funções não exportadas normalmente.

Quando temos funções exportadas, é uma boa prática ter um comentário acima explicando o que elas fazem:

```go
// Somar soma dois números.
func Somar() {
    // ...
}
```

Em [pkg.go.dev](https://pkg.go.dev/) há a documentação dos pacotes da linguagem Go, incluindo os pacotes da biblioteca padrão.

## Pacote externo

Para instalar/adicionar um pacote externo, podemos usar:

```bash
go get [url_do_pacote]
```

O arquivo `go.mod` registra as dependências diretas do projeto.

O arquivo `go.sum` armazena os **checksums (hashes)** das dependências utilizadas, ajudando o Go a verificar a integridade das versões dos módulos baixados.

Para remover dependências que não estão sendo utilizadas e organizar as dependências do projeto, usamos:

```bash
go mod tidy
```

Para utilizar o pacote no código, precisamos importá-lo, normalmente utilizando o caminho do pacote:

```go
import "url_do_pacote"
```

Depois, podemos chamar as funções exportadas desse pacote:

```go
pacote.Funcao()
```

---

## Material de leitura

- [Golang — Victor Stein](https://victorstein.dev/posts/golang/012-golang/)
- [Módulos em Go](https://micilini.com/conteudos/golang/modulos-em-go)
- [Go Modules Reference](https://go.dev/ref/mod/)