package main

import "fmt"

func f1(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

func f2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * 2
		}
	}()

	return out
}

func f3(in <-chan int) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for num := range in {
			fmt.Println(num)
		}
	}()

	return done
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	ch1 := f1(nums)
	ch2 := f2(ch1)
	done := f3(ch2)
	<-done
}
