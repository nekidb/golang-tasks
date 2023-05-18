package main

import "fmt"

// функция отправляет числа из nums в канал out
func f1(nums []int, out chan<- int) {
	for _, num := range nums {
		out <- num
	}
	close(out)
}

// функция читает числа из канала in, умножает их на 2 и отправляет в канал out
func f2(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

// функция читает числа из канала in и выводит в stdout
// по завершению выполнения отправляем в канал done
func f3(done chan<- struct{}, in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
	done <- struct{}{}
}

func main() {
	// запускаем 3 горутины:
	//  первая пишет в канал ch1
	//  вторая читает из канала ch1 и отправляет в ch2 удвоенное число
	//  третья читает из ch2 и выводит в stdout
	// чтобы горутина main не завершилась, читаем из канала done
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go f1(nums, ch1)

	go f2(ch1, ch2)

	done := make(chan struct{})
	go f3(done, ch2)

	<-done
}
