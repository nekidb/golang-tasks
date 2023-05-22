package main

import (
	"fmt"
	"math/big"
)

func main() {
	// создаем два числа больше 2^20
	a := new(big.Int)
	a.SetString("555_555_555_555_555_555_555_555", 0)
	b := new(big.Int)
	b.SetString("222_222_222_222_222_222_222_222", 0)

	// находим сумму
	sum := new(big.Int)
	sum.Add(a, b)
	// находим разность
	sub := new(big.Int)
	sub.Sub(a, b)
	// находим деление
	div := new(big.Int)
	div.Div(a, b)
	// находим произведение
	mul := new(big.Int)
	mul.Mul(a, b)

	fmt.Println("a:   ", a)
	fmt.Println("b:   ", b)
	fmt.Println("Sum: ", sum)
	fmt.Println("Sub: ", sub)
	fmt.Println("Div: ", div)
	fmt.Println("Mul: ", mul)

}
