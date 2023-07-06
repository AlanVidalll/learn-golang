package main

import "fmt"

// existem esses exemplos de switch
func weekDay(number int) string {
	switch number {
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
		return "Sabado"
	default:
		return "Número inválido"
	}
}

func weekDay2(number int) string {
	switch {
	case number == 1:
		return "Domingo"
	case number == 2:
		return "Segunda-feira"
	case number == 3:
		return "Terça-feira"
	case number == 4:
		return "Quarta-feira"
	case number == 5:
		return "Quinta-feira"
	case number == 6:
		return "Sexta-feira"
	case number == 7:
		return "Sabado"
	default:
		return "Nuúmero inválido"
	}
}

func weekDay3(number int) string {
	var day3 string
	switch {
	case number == 1:
		day3 = "Domingo"
	case number == 2:
		day3 = "Segunda-feira"
	case number == 3:
		day3 = "Terça-feira"
	case number == 4:
		day3 = "Quarta-feira"
	case number == 5:
		day3 = "Quinta-feira"
	case number == 6:
		day3 = "Sexta-feira"
	case number == 7:
		day3 = "Sabado"
	default:
		day3 = "Nuúmero inválido"
	}

	return day3
}

func main() {
	day := weekDay(4)
	day2 := weekDay2(4)
	day3 := weekDay3(6)

	fmt.Println(day)
	fmt.Println(day2)
	fmt.Println(day3)
}
