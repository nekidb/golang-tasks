package main

import (
	"fmt"
	"sync"
)

func square(n int) int {
	return n * n
}

func squaresSender(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		var wg sync.WaitGroup
		wg.Add(len(nums))

		for _, num := range nums {
			go func(n int) {
				wg.Done()

				out <- square(n)
			}(num)
		}

		wg.Wait()
	}()

	return out
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	squared := squaresSender(nums)

	var sum int
	for s := range squared {
		sum += s
	}

	fmt.Println(sum)
}
