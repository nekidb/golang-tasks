package main

import "fmt"

// Когда бит нужно поменять на 0, мы будем применять f1. Когда на 1, тогда f2
// x | y | f1 | f2|
// 0 | 0 | 0  | 0 |
// 0 | 1 | 0  | 1 |
// 1 | 0 | 1  | 1 |
// 1 | 1 | 0  | 1 |

// Получаем формулы:
// f1 = X & ^Y
// f2 = (^X & Y) || (X & ^Y) || (X & Y)

// функция меняет в числе num бит с номером bit на значение val
func changeBit(num, bit, val int64) int64 {
	m := int64(1 << bit)
	if val == 0 {
		return num & ^m
	}
	return (^num & m) | (num & ^m) | (num & m)
}

func main() {
	number := int64(123456789)

	fmt.Printf("Исходное число:       %b\n", number)
	fmt.Printf("Меняем 3-ый бит на 1: %b\n", changeBit(number, 3, 1))
	fmt.Printf("Меняем 2-ый бит на 0: %b\n", changeBit(number, 2, 0))
}
