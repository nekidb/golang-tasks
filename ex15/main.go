package main

var justString string

func someFunc() {
	// вместо того, чтобы записывать результат функции createHugeString в промежуточную переменную
	// мы записываем "обрезанный" результат сразу в переменную justString
	justString = createHugeString()[:100]
}

func main() {
	someFunc()
}
