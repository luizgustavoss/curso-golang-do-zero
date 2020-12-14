package main

import "fmt"


func main()  {

	soma := somar(10, 34)
	f(soma)

	r1, r2 := calculos(34, 20)
	f(r1, r2)

	rn1, rn2 := retornoNomeado(50, 30)
	f(rn1, rn2)

	total := funcaoVariatica(2, 50, 43, 23, 9)
	f(total)

	retornoFuncAnon := func(num int32) int32 {
		fmt.Println("Função Anônima")
		return num
	}(50)

	fmt.Printf("Retorno da função anônima = %d ", retornoFuncAnon)
}


func funcaoVariatica(numeros...int) int {

	total := 0

	for _, numero := range numeros{
		total += numero
	}

	return total
}

func retornoNomeado(n1, n2 int) (soma int, subtracao int){
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func somar(n1, n2 int8) int8 {
	return n1 + n2;
}

func calculos(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

var f = func(a ...interface{}){
	fmt.Println(a)
}
