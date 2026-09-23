package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите текст: ")
	scanner.Scan()
	text := strings.ToLower(scanner.Text())

	words := strings.Fields(text)

	counter := make(map[string]int)
	for _, word := range words {
		counter[word]++
	}

	keys := make([]string, 0, len(counter))
	for word := range counter {
		keys = append(keys, word)
	}
	sort.Strings(keys)

	fmt.Println("\nЧастота слов (по алфавиту): ")
	for _, word := range keys {
		fmt.Printf(" %-15s %d\n", word, counter[word])
	}
	fmt.Printf("\nВсего уникальных слов: %d\n", len(counter))
	fmt.Printf("Всего слов: %d\n", len(words))
}
