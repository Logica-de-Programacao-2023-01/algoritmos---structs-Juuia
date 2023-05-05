package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
	ano    int
}

func criaLivro(titulo string, autor string, ano int) Livro {
	return Livro{titulo: titulo, autor: autor, ano: ano}
}

func main() {
	l := criaLivro("Julia a bela,", "Julia,", 2023)
	fmt.Println(l)
}
