package main

import "fmt"

func main() {

	var var1 int32
	var varPointer *int32

	var1 = 10
	varPointer = &var1

	fmt.Println(var1, *varPointer)

	var1 = 15
	fmt.Println(var1, *varPointer)

	*varPointer = 23

	fmt.Println(var1, *varPointer)
}
