package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")
	for {
	userKg, userHeight := getUserInput()
	IMT, err := calculateIMT(&userKg, userHeight)
	if err != nil {
		fmt.Println("Не заданны параметры для расчета")
		continue
	}
	outputResult(IMT)
	isRepeateCalculation := checkRepeaCalculation()
	if !isRepeateCalculation {
		break
	}
}
func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш идекс массы тела: %.0f", imt)
	fmt.Print(result)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
		case calculateIMT() < 18.5:
		fmt.Println("У вас дефицит массы тела")
		case imt < 25:
		fmt.Println("У вас нормальный вес")
		case imt < 30:
		fmt.Println("У вас избыточный вес")
		default:
			fmt.Println("У вас степень ожирения")
		
	}
}
func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if eserKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}
func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}

func checkRepeaCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчет (y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}