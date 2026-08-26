package main

import (
	"fmt"
	"strings"
)

const (
	usdInEuro = 0.86
	usdInRub  = 84.38
)

func main() {
	fmt.Println(" КОНВЕРТЕР ВАЛЮТ ")
	fmt.Printf("Курсы: 1 USD = %.2f EUR, 1 USD = %.2f RUB\n", usdInEuro, usdInRub)
	fmt.Println()

	from := inputCurrency("Введите исходную валюту")
	amount := inputAmount()
	to := inputCurrency("Введите целевую валюту")

	result := convertCurrency(amount, from, to)

	fmt.Println()
	fmt.Println(" РЕЗУЛЬТАТ ")
	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
}

func isValidCurrency(currency string) bool {
	switch currency {
	case "USD", "EUR", "RUB":
		return true
	default:
		return false
	}
}

func inputCurrency(prompt string) string {
	var currency string

	for {
		fmt.Printf("%s (USD, EUR, RUB): ", prompt)
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		if isValidCurrency(currency) {
			return currency
		}
		fmt.Println("ОШИБКА! Введите USD, EUR или RUB")
	}
}

func inputAmount() float64 {
	var amount float64

	for {
		fmt.Print("Введите сумму: ")
		_, err := fmt.Scan(&amount)
		if err == nil && amount > 0 {
			return amount
		}
		fmt.Println("ОШИБКА! Введите положительное число")
	}
}

func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	if fromCurrency == toCurrency {
		return amount
	}

	var amountInUSD float64

	switch fromCurrency {
	case "USD":
		amountInUSD = amount
	case "EUR":
		amountInUSD = amount / usdInEuro
	case "RUB":
		amountInUSD = amount / usdInRub
	}

	switch toCurrency {
	case "USD":
		return amountInUSD
	case "EUR":
		return amountInUSD * usdInEuro
	case "RUB":
		return amountInUSD * usdInRub
	}

	return 0
}
