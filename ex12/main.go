package main

import (
	"fmt"
	"strings"
)

type Set map[string]struct{}

func (s Set) Add(val string) {
	s[val] = struct{}{}
}

func (s Set) String() string {
	out := strings.Builder{}

	out.WriteString("{ ")
	for k, _ := range s {
		out.WriteString(k)
		out.WriteString(" ")
	}
	out.WriteString("}")

	return out.String()
}

func main() {
	someStrings := []string{"cat", "cat", "dog", "cat", "tree"}

	s := Set{}
	for _, str := range someStrings {
		s.Add(str)
	}

	fmt.Println(s)
}
