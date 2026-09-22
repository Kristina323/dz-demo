package main

import "fmt"

func main() {
	phonebook := map[string]string{
		"Иван":      "89181324567",
		"Анастасия": "89004556892",
		"Роман":     "89384005675",
	}

	var search string
	fmt.Print("Введите имя: ")
	fmt.Scan(&search)

	if phone, exists := phonebook[search]; exists {
		fmt.Printf("%s: %s\n", search, phone)
	} else {
		fmt.Println("Контакт не найден")
	}
}
