package main

import "fmt"

func log(obj interface{}){
	fmt.Println(obj)
}

func main() {

	log("Luiz")
	log("Mariana")

	slc := []int32{23, 45, 56, 76, 4567}
	log(slc)
}
