package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secret := rand.Intn(100) + 1
	maxAttempts := 7
	attempts := 0

	fmt.Println("Я загадал число от 1 до 100.")

	for attempts < maxAttempts {
		attempts++
		fmt.Printf("Попытка %d/%d. Введите число:", attempts, maxAttempts)

		var guess int
		fmt.Scan(&guess)

		if guess < 1 || guess > 100 {
			fmt.Println("Число должно быть от 1 до 100!")
			attempts--
			continue
		}

		if guess == secret {
			fmt.Printf("Правильно! Ты угадал за %d попыток!\n", attempts)
			break
		} else if guess < secret {
			fmt.Printf("Загаданное число БОЛЬШЕ.")
		} else {
			fmt.Println("Загаданное число МЕНЬШЕ.")
		}
	}

	fmt.Printf("Ты проиграл. Загаданное число было %d./n", secret)
}
