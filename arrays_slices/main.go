package main

import "fmt"

func main() {

	var array1 [5]int32
	array1[0] = 22
	fmt.Println(array1)

	array2 := [2]string{"Luiz", "Gustavo"}
	fmt.Println(array2)

	array3 := [...]string{"Luiz", "Gustavo", "Stábile", "Souza"}
	fmt.Println(array3)

	slice := []int32{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(slice)

	slice = append(slice, 23)
	fmt.Println(slice)

	slice2 := array3[0:3]
	array3[2] = "Zambas"

	fmt.Println(slice2)

	slice4 := make([]string, 5)
	slice4[2] = "Luiz"
	fmt.Println(slice4)
}
