package main

import "fmt"

func main() {
	var weight float64
	var height float64

	fmt.Print("Введите свой вес (кг): ")
	fmt.Scan(&weight)

	fmt.Print("Введите свой рост (м): ")
	fmt.Scan(&height)

	bmi := weight / (height * height)
	fmt.Printf("Твой ИМТ: %.1f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("У вас недостаточный вес!")
	} else if bmi < 25 {
		fmt.Println("У вас нормальный вес!")
	} else if bmi < 30 {
		fmt.Println("У вас избыточный вес!")
	} else {
		fmt.Println("У вас ожирение!")
	}
}
