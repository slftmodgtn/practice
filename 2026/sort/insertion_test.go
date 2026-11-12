package sort

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	testSortTable(t, []testCase[int]{
		{"nil", []int(nil), []int(nil)},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"small", []int{2, 1}, []int{1, 2}},
		{"same", []int{1, 1}, []int{1, 1}},
		{"sorted", []int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"reversed", []int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
		{"doubled", []int{3, 3, 2, 2, 1, 10, 0}, []int{0, 1, 2, 2, 3, 3, 10}},
		{"random", []int{10, 6, 2, 0, 34, 107, 36}, []int{0, 2, 6, 10, 34, 36, 107}},
	})
}

func TestInsertionSortFloat64(t *testing.T) {
	testSortTable(t, []testCase[float64]{
		{"reversed", []float64{5.5, 4.4, 3.3, 2.2, 1.1}, []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"same", []float64{1.5, 1.5, 1.5}, []float64{1.5, 1.5, 1.5}},
		{"random", []float64{10.5, -3.25, 0, 7.75, -100.5}, []float64{-100.5, -3.25, 0, 7.75, 10.5}},
	})
}

func TestInsertionSortString(t *testing.T) {
	testSortTable(t, []testCase[string]{
		{"reversed", []string{"delta", "charlie", "bravo", "alpha"}, []string{"alpha", "bravo", "charlie", "delta"}},
		{"same", []string{"x", "x"}, []string{"x", "x"}},
		{"case", []string{"banana", "Apple", "cherry"}, []string{"Apple", "banana", "cherry"}},
	})
}

func TestInsertionSortRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	t.Run("int", func(t *testing.T) { testSortRandom(t, r, randomInts) })
	t.Run("float64", func(t *testing.T) { testSortRandom(t, r, randomFloats) })
	t.Run("string", func(t *testing.T) { testSortRandom(t, r, randomStrings) })
}

func FuzzInsertionSort(f *testing.F) {
	f.Add([]byte{3, 1, 2}, "cba")
	f.Add([]byte{}, "")
	f.Fuzz(func(t *testing.T, data []byte, s string) {
		arr := make([]int, len(data))
		for i, b := range data {
			arr[i] = int(b) - 128
		}
		want := slices.Clone(arr)
		slices.Sort(want)
		for _, sc := range sortCases[int]() {
			got := slices.Clone(arr)
			sc.fn(got)
			if !slices.Equal(got, want) {
				t.Fatalf("%s(%v) = %v, want %v", sc.name, arr, got, want)
			}
		}

		strs := strings.Split(s, ",")
		wantStrs := slices.Clone(strs)
		slices.Sort(wantStrs)
		for _, sc := range sortCases[string]() {
			got := slices.Clone(strs)
			sc.fn(got)
			if !slices.Equal(got, wantStrs) {
				t.Fatalf("%s(%q) = %q, want %q", sc.name, strs, got, wantStrs)
			}
		}
	})
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

