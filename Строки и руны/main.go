package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var input string
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	lowercase := strings.ToLower(input)
	fmt.Println("Строка в нижнем регистре:", lowercase)

	words := strings.Fields(input)
	fmt.Println("Количество слов:", len(words))

	if strings.Contains(lowercase, "go") {
		fmt.Println("Содержит слово \"go\"")
	} else {
		fmt.Println("Не содержит слово \"go\"")
	}
}
