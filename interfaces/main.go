package main

import (
	"fmt"
	"math"
)

type forma interface{
	area() float64
	nome() string
}

func escreverArea(f forma){
	fmt.Printf("A área da forma %s é %0.2f \n", f.nome(), f.area())
}

type retangulo struct{
	altura float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (r retangulo) nome() string {
	return "Retângulo"
}

type circulo struct{
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (c circulo) nome() string {
	return "Circulo"
}

func main() {

	circulo := circulo{raio: 34}
	escreverArea(circulo)

	retangulo := retangulo{altura: 23, largura: 25}
	escreverArea(retangulo)
}
