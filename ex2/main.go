package main

import (
	"fmt"
	"sync"
)

// Функция вычисляет квадрат числа и выводит в stdout
func printSquare(x int) {
	fmt.Printf("Square of %d:\t%d\n", x, x*x)
}

func main() {
	nums := [5]int{2, 4, 6, 8, 10}

	// wg будет ожидать завершения набора горутин
	wg := sync.WaitGroup{}

	// записываем в wg количество горутин, которые будут созданы
	wg.Add(len(nums))
	for _, num := range nums {
		go func(n int) {
			// по окончанию выполнения функции, указываем в wg что горутина больше неактивна
			defer wg.Done()
			printSquare(n)
		}(num)
	}

	// ожидаем завершения горутин
	wg.Wait()
}
