package main

var justString string

func someFunc() string {
	return createHugeString()[:100]
}

func main() {
	justString = someFunc()

}
