package main

import (
	"fmt"
	"strings"
)

// функция переворачивает слова в строке
func revertWords(s string) string {
	// разбиваем строку на слова
	words := strings.Split(s, " ")

	result := strings.Builder{}
	// начиная с конца массива слов добавляем каждое слово в билдер
	for i := len(words) - 1; i >= 0; i-- {
		result.WriteString(words[i])
		if i != 0 {
			result.WriteString(" ")
		}
	}

	return result.String()
}

func main() {
	s := "snow dog sun"

	fmt.Println(revertWords(s))
}
