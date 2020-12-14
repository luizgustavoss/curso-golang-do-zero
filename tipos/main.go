package main

import (
	"errors"
	"fmt"
)

func main()  {

	var i8 int8 = -8
	var i16 int16 = -16
	var i32 int32 = -32
	var i64 int64 = -64
	fmt.Println(i8, i16, i32, i64)

	var ui8 uint8 = 8
	var ui16 uint16 = 16
	var ui32 uint32 = 32
	var ui64 uint64 = 64
	fmt.Println(ui8, ui16, ui32, ui64)

	// int32 = rune
	var rune32 rune = 32
	fmt.Println(rune32)

	// byte = int8
	var byte8 byte = 128
	fmt.Println(byte8)

	var f32 float32 = 12345.653
	fmt.Println(f32)

	var f64 float64 = 2345.342345
	fmt.Println(f64)

	var nome string = "Luiz"
	fmt.Println(nome)

	var sobrenome string = "Stábile"
	fmt.Println(sobrenome)

	var error error = errors.New("Exemplo de erro!")
	fmt.Println(error)
}
