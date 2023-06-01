package main

import "fmt"

func main() {
	x, y := 10, 15

	fmt.Printf("До перестановки\nx: %d\ny: %d\n", x, y)

	x, y = y, x

	fmt.Printf("После перестановки\nx: %d\ny: %d", x, y)
}
