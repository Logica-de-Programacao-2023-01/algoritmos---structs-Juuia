package main

import "fmt"

type Retangulo struct {
	largura float64
	altura  float64
}

func areaRetangulo(ret Retangulo) float64 {
	return ret.largura * ret.altura
}

func main() {
	var ret Retangulo

	fmt.Print("Altura: ")
	fmt.Scan(&ret.altura)
	fmt.Print("Largura: ")
	fmt.Scan(&ret.largura)
	a := areaRetangulo(ret)
	fmt.Println("A área é de: ", a)

}
