package main

import "fmt"

type Human struct {
	Height int
	Age    int
}

// Для встраивания структуры Human в структуру Action, указываем Human как поле, но без имени.
type Action struct {
	Human
	ActionField string
}

func main() {
	fmt.Println("Hello")
}
