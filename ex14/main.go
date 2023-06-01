package main

import (
	"fmt"
	"reflect"
	"strings"
)

// функция выводит тип аргумента v
func printTypeOf(v any) {
	// определяем тип переменной v
	// если тип состоит из двух слов, то вытаскиваем первой (для типа chan)
	t := strings.Split(reflect.TypeOf(v).String(), " ")[0]

	switch t {
	case "int":
		fmt.Println("Это тип int")
	case "string":
		fmt.Println("Это тип string")
	case "bool":
		fmt.Println("Это тип bool")
	case "chan":
		fmt.Println("Это тип channel")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	// создаем переменные разных типов и выводим их тип
	v1 := 10
	v2 := "string"
	v3 := true
	v4 := make(chan int)

	printTypeOf(v1)
	printTypeOf(v2)
	printTypeOf(v3)
	printTypeOf(v4)
}
