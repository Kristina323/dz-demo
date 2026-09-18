package main

import "fmt"

func main() {
	var day int
	fmt.Print("Введите номер дня(1 - 7): ")
	fmt.Scan(&day)

	var name string
	switch day {
	case 1:
		name = "Понедельник"
	case 2:
		name = "Вторник"
	case 3:
		name = "Среда"
	case 4:
		name = "Четверг"
	case 5:
		name = "Пятница"
	case 6:
		name = "Суббота"
	case 7:
		name = "Воскресенье"
	default:
		fmt.Print("Неверный день:")
		return
	}
	fmt.Println("День недели:", name)

	if day < 1 || day > 7 {
		fmt.Println("Неверный день недели:")
	} else if day < 5 {
		fmt.Println("Будний день")
	} else {
		fmt.Println("Выходной день")
	}
}
