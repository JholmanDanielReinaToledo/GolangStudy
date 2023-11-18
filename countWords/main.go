package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./lorem ipsum.txt")
	if err != nil {
		fmt.Println("Ha ocurrido un error")
		return
	}

	text := strings.Split(cleanText(string(data)), " ")

	contador := make(map[string]int)

	for _, word := range text {
		word := strings.ToLower(word)
		contador[word]++
	}

	for word, count := range contador {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func cleanText(text string) string {
	cleanString1 := strings.ReplaceAll(strings.ReplaceAll(text, ".", " "), ",", " ")

	return strings.ReplaceAll(cleanString1, "\n", " ")
}
