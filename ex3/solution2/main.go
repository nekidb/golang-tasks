package main

import (
	"fmt"
)

// функция считает квадрат числа
func square(n int) int {
	return n * n
}

// функция вычисляет квадрат числа n и отправляет в канал out
func squareAndSend(n int, out chan<- int) {
	out <- square(n)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	squared := make(chan int, len(nums))

	// для каждого числа из num в отдельной горутине выполняем функцию squareAndSend
	for _, num := range nums {
		go squareAndSend(num, squared)
	}

	// читаем из канала и записываем результат в sum
	// когда канал закроется, тогда завершится цикл
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += <-squared
	}

	fmt.Println(sum)
}
