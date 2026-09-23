package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var operation string
	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	fmt.Scan(&operation)

	var numbersInput string
	fmt.Print("Введите числа через запятую (например: 1,2,3): ")
	fmt.Scan(&numbersInput)

	parts := strings.Split(numbersInput, ",")
	numbers := []float64{}

	for _, part := range parts {
		part = strings.TrimSpace(part)
		num, _ := strconv.ParseFloat(part, 64)
		numbers = append(numbers, num)
	}

	fmt.Println("Числа:", numbers)

	switch operation {
	case "SUM":
		sum := 0.0
		for _, num := range numbers {
			sum += num
		}
		fmt.Println("Сумма =", sum)

	case "AVG":
		sum := 0.0
		for _, num := range numbers {
			sum += num
		}
		fmt.Println("Среднее =", sum/float64(len(numbers)))

	case "MED":
		for i := 0; i < len(numbers); i++ {
			for j := i + 1; j < len(numbers); j++ {
				if numbers[i] > numbers[j] {
					numbers[i], numbers[j] = numbers[j], numbers[i]
				}
			}
		}

		n := len(numbers)
		if n%2 == 0 {
			fmt.Println("Медиана =", (numbers[n/2-1]+numbers[n/2])/2)
		} else {
			fmt.Println("Медиана =", numbers[n/2])
		}

	default:
		fmt.Println("Неверная операция! Используйте AVG, SUM или MED")
	}
}
