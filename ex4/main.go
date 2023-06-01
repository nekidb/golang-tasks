package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func printValue(id int, val int) {
	fmt.Printf("Worker %d: %d\n", id, val)
}

func main() {
	n := flag.Int("n", 5, "-n")
	flag.Parse()

	pool := make(chan int, *n)
	defer close(pool)
	for i := 1; i <= *n; i++ {
		pool <- i
	}

	ch := make(chan int)
	defer close(ch)

	cancel := make(chan struct{})

	go func() {
		for {
			id := <-pool
			go func() {
				select {
				case <-cancel:
					return
				default:
					printValue(id, <-ch)
				}
				pool <- id
			}()
		}
	}()

	idx, maxIdx := 0, 50
	for {
		ch <- rand.Int()
		time.Sleep(100 * time.Millisecond)
		if idx > maxIdx {
			close(cancel)
			break
		}
		idx++
	}
}
