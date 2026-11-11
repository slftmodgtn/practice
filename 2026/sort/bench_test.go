package sort

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
)

func benchmarkData(size int, kind string) []int {
	r := rand.New(rand.NewPCG(uint64(size), 1))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = r.IntN(size * 10)
	}
	switch kind {
	case "sorted":
		slices.Sort(arr)
	case "reversed":
		slices.Sort(arr)
		slices.Reverse(arr)
	}
	return arr
}

func BenchmarkInsertionSort(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	kinds := []string{"random", "sorted", "reversed"}
	impls := append(sortCases[int](), sortImpl[int]{"stdlib", func(arr []int) { slices.Sort(arr) }})

	for _, size := range sizes {
		for _, kind := range kinds {
			src := benchmarkData(size, kind)
			for _, impl := range impls {
				b.Run(fmt.Sprintf("%s/%s/n=%d", impl.name, kind, size), func(b *testing.B) {
					b.ReportAllocs()
					arr := make([]int, len(src))
					for b.Loop() {
						copy(arr, src) // Работает без новых аллокаций памяти!
						impl.fn(arr)
					}
				})
			}
		}
	}
}
