package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Set map[int]struct{}

func (s Set) Add(val int) {
	s[val] = struct{}{}
}

// функция находит и возвращает пересечение c множеством
func (s Set) Intersection(s2 Set) Set {
	out := Set{}
	// проходимся по элементам первого множества и проверяем, есть ли такой же элемент во втором
	// если элемент найден, то добавляем его в out
	for k, _ := range s {
		if _, ok := s2[k]; ok {
			out[k] = struct{}{}
		}
	}

	return out
}

func (s Set) String() string {
	out := strings.Builder{}

	out.WriteString("{ ")
	for k, _ := range s {
		out.WriteString(strconv.Itoa(k))
		out.WriteString(" ")
	}
	out.WriteString("}")

	return out.String()
}

func main() {
	// для поиска пересечения используем метод Intersection типа Set
	s1 := Set{}
	s1.Add(10)
	s1.Add(20)
	s1.Add(30)

	s2 := Set{}
	s2.Add(20)
	s2.Add(30)
	s2.Add(40)

	s3 := s2.Intersection(s1)

	fmt.Println("First set: ", s1)
	fmt.Println("Second set: ", s2)
	fmt.Println("Intersection: ", s3)
}

//
// type Set struct {
// 	data map[int]int
// }
//
// func NewSet() Set {
// 	return Set{data: make(map[int]int)}
// }
//
// func (s *Set) Add(v int) {
// 	s.data[v] = s.data[v] + 1
// }
//
// func (s *Set) Count(k int) int {
// 	v, ok := s.data[k]
// 	if !ok {
// 		return 0
// 	}
// 	return v
// }
//
// func (s *Set) String() string {
// 	out := strings.Builder{}
//
// 	out.WriteString("{ ")
// 	for k, _ := range s.data {
// 		out.WriteString(strconv.Itoa(k))
// 		out.WriteString(" ")
// 	}
// 	out.WriteString("}")
//
// 	return out.String()
// }
//
// func main() {
// 	s1 := NewSet()
//
// 	s1.Add(1)
// 	// s1.Add(2)
// 	// s1.Add(2)
//
// 	s2 := NewSet()
//
// 	s2.Add(1)
// 	// s2.Add(2)
//
// 	// s3 := NewSet()
// 	// for k, v := range s1 {
// 	// 	if
// 	// }
//
// 	fmt.Println(s1.String())
// 	fmt.Println(s2.String())
// }
