package main

import (
	"fmt"
	"strings"
)

// функция проверяет, что все символы в строке уникальны
func IsUnique(s string) bool {
	// для регистронезависимости приводим строку к нижнему регистру
	s = strings.ToLower(s)
	// создаем мапу, которая будет содержать вхождения символов в строке
	chars := make(map[rune]struct{})
	// проходимся по каждому символу и проверяем, присутствует ли он в мапе
	//  если да, то возвращаем false, иначе добавляем этот символ в мапу
	for _, r := range s {
		if _, ok := chars[r]; ok {
			return false
		}
		chars[r] = struct{}{}
	}
	return true
}

func main() {
	examples := []string{"abcd", "abCdefA", "aabcd"}

	for _, e := range examples {
		fmt.Printf("%s - %v\n", e, IsUnique(e))
	}

}
