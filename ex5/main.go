package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// функция создает канал в который бесконечно отправляет рандомные целые числа
func writer(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case out <- rand.Int():
				continue
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

// функция бесконечно читает из канала числа
func reader(ctx context.Context, in <-chan int) {
	go func() {
		for {
			select {
			case val := <-in:
				fmt.Println(val)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func main() {
	// Одна горутина бесконечно пишет в канал, другая бесконечно читает из него.
	// По истечении N секунд программа завершается
	N := 1

	dur := time.Duration(N) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), dur)
	defer cancel()

	ch := writer(ctx)
	reader(ctx, ch)

	// Ждем достаточно, чтобы выполнился таймаут
	time.Sleep(time.Duration(N+1) * time.Second)
}
