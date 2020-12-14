package main

import "fmt"

func main()  {

	var variavel1 string = "Variavel 1"
	variavel2 := "Variável 2"

	var (
		variavel3 string = "Variavel 3"
		variavel4 = "Variavel 4"
	)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	const contante1 = "Constante 1"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)
	fmt.Println(contante1)
}
