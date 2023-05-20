package main

import "testing"

func BenchmarkRemove1(b *testing.B) {
	nums := make([]int, 1000000)

	for i := 0; i < b.N; i++ {
		Remove1(nums, len(nums)/2)
	}
}

func BenchmarkRemove2(b *testing.B) {
	nums := make([]int, 1000000)

	for i := 0; i < b.N; i++ {
		Remove2(nums, len(nums)/2)
	}
}
