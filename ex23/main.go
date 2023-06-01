package main

import "fmt"

// функция удаляет элемент с номером idx из слайса путем смещения правой части на один элемент влево
func Remove1(nums []int, idx int) []int {
	for i := idx; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}

	nums = nums[:len(nums)-1]
	return nums
}

// функция удаляет элемент путем склеивания двух частей исходного слайса, которые не содержат idx-ый элемент
func Remove2(nums []int, idx int) []int {
	return append(nums[:idx], nums[idx+1:]...)
}

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("До удаления: ", x)
	x = Remove1(x, 0)
	fmt.Println("После удаления: ", x)

	x = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("До удаления: ", x)
	x = Remove2(x, 0)
	fmt.Println("После удаления: ", x)
}
