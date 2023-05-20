package main

import (
	"fmt"
	"unicode/utf8"
)

// функция переворачивает строку
func revert(s string) string {
	// создаем слайс рун, который будет содержать перевернутую строку
	runes := make([]rune, utf8.RuneCountInString(s))
	idx := 1
	// последовательно берем символы строки и с конца заполняем массив runes
	for _, r := range s {
		runes[len(runes)-idx] = r
		idx++
	}

	return string(runes)
}

func main() {
	s := "Ёлка"

	fmt.Println(revert(s))
}
