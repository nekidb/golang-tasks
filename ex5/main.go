package main

import (
	"math/rand"
	"time"
)

// функция бесконечно отправляет рандомные целые числа в канал out
func writeToChan(out chan<- int) {
	for {
		out <- rand.Int()
	}
}

// функция бесконечно читает из канала in целые числа
func readFromChan(in <-chan int) {
	for {
		<-in
	}
}

func main() {
	N := 5
	ch := make(chan int)

	// создаем горутину, которая отправляет значения в канал
	go writeToChan(ch)

	// создаем горутину, которая читает значения из канала
	go readFromChan(ch)

	// ждем N секунд и завершаем программу
	time.Sleep(time.Duration(N) * time.Second)
}
