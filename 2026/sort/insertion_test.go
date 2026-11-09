package main

import (
	"math/rand/v2"
	"slices"
	"testing"
)

type testCase struct {
	name string
	in   []int
	want []int
}

func sortCases() []struct {
	name string
	fn   func([]int)
} {
	return []struct {
		name string
		fn   func([]int)
	}{
		{"classic", InsertionSortClassic},
		{"range", InsertionSortRange},
	}
}

func getTestData() []testCase {
	return []testCase{
		{"nil", []int(nil), []int(nil)},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"two", []int{1, 2}, []int{1, 2}},
		{"small", []int{2, 1}, []int{1, 2}},
		{"same", []int{1, 1}, []int{1, 1}},
		{"sorted", []int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"reversed", []int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
		{"doubled", []int{3, 3, 2, 2, 1, 10, 0}, []int{0, 1, 2, 2, 3, 3, 10}},
		{"random", []int{10, 6, 2, 0, 34, 107, 36}, []int{0, 2, 6, 10, 34, 36, 107}},
	}
}

func TestInsertionSort(t *testing.T) {
	for _, sc := range sortCases() {
		for _, tt := range getTestData() {
			t.Run(sc.name+"/"+tt.name, func(t *testing.T) {
				got := slices.Clone(tt.in)
				sc.fn(got)
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", sc.name, tt.in, got, tt.want)
				}
			})
		}
	}
}

func TestInsertionSortRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	for _, sc := range sortCases() {
		t.Run(sc.name, func(t *testing.T) {
			for range 100 {
				arr := make([]int, r.IntN(100))
				for i := range arr {
					arr[i] = r.IntN(201) - 100
				}
				want := slices.Clone(arr)
				slices.Sort(want)
				orig := slices.Clone(arr)

				sc.fn(arr)
				if !slices.Equal(arr, want) {
					t.Fatalf("%s(%v) = %v, want %v", sc.name, orig, arr, want)
				}
			}
		})
	}
}

func FuzzInsertionSort(f *testing.F) {
	f.Add([]byte{3, 1, 2})
	f.Add([]byte{})
	f.Fuzz(func(t *testing.T, data []byte) {
		arr := make([]int, len(data))
		for i, b := range data {
			arr[i] = int(b) - 128
		}
		want := slices.Clone(arr)
		slices.Sort(want)
		for _, sc := range sortCases() {
			got := slices.Clone(arr)
			sc.fn(got)
			if !slices.Equal(got, want) {
				t.Fatalf("%s(%v) = %v, want %v", sc.name, arr, got, want)
			}
		}
	})
}

func BenchmarkInsertionSortClassic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSortClassic([]int{10, 6, 2, 0, 34, 107, 36})
	}
}

func BenchmarkInsertionSortRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSortRange([]int{10, 6, 2, 0, 34, 107, 36})
	}
}
