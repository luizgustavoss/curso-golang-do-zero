package main

import (
	"fmt"
)

func main() {

	numero := 10

	// if-else simples
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// if-init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("outro número é maior que zero")
	}

	sabado := diaDaSemana(7)
	fmt.Println(sabado)

	invalido := diaDaSemana(200)
	fmt.Println(invalido)

	saturday := dayOfWeek(7)
	fmt.Println(saturday)

	invalid := dayOfWeek(200)
	fmt.Println(invalid)

}


func diaDaSemana(dia int) string {

	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func dayOfWeek(day int) string {

	switch {
	case day == 1:
		return "Sunday"
	case day == 2:
		return "Monday"
	case day == 3:
		return "Thursday"
	case day == 4:
		return "Wednesday"
	case day == 5:
		return "Tuesday"
	case day == 6:
		return "Friday"
	case day == 7:
		return "Saturday"
	default:
		return "Invalid day"
	}
}