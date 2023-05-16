package main

import (
	"fmt"
	"sync"
)

// функция считает квадрат числа
func square(n int) int {
	return n * n
}

// функция конкурентно вычисляет квадраты чисел из nums и записывае их в канал out
func squaresSender(nums []int, out chan<- int) {
	// создаем wg, чтобы можно было понять когда завершаться все горутины, вычисляющие квадраты
	wg := sync.WaitGroup{}

	// для каждого числа из num создаем горутину, которая вычисляет квадрат числа и записывает
	// результат в канал out
	for _, num := range nums {
		// перед созданием горутины, увеличиваем счетчик в wg
		wg.Add(1)
		go func(n int) {
			// по окончанию выполнения функции уменьшаем счетчик в wg
			wg.Done()
			// передаем в канал квадрат числа
			out <- square(n)
		}(num)
	}
	// ожидаем завершения всех горутин
	wg.Wait()
	// закрываем канал
	close(out)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	squared := make(chan int)

	go squaresSender(nums, squared)

	sum := 0
	// читаем из канала и записываем результат в sum
	// когда канал закроется, тогда завершится цикл
	for s := range squared {
		sum += s
	}

	fmt.Println(sum)
}
