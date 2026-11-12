package sort

import (
	"cmp"
	"math/rand/v2"
	"slices"
	"testing"
)

type testCase[T cmp.Ordered] struct {
	name string
	in   []T
	want []T
}

type sortImpl[T cmp.Ordered] struct {
	name string
	fn   func([]T)
}

func sortCases[T cmp.Ordered]() []sortImpl[T] {
	return []sortImpl[T]{
		{"classic", InsertionSortClassic[T]},
		{"range", InsertionSortRange[T]},
		{"custom-binary", InsertionSortCustomBinary[T]},
		{"std-binary", InsertionSortStdBinary[T]},
	}
}

func testSortTable[T cmp.Ordered](t *testing.T, cases []testCase[T]) {
	for _, sc := range sortCases[T]() {
		for _, tt := range cases {
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

func randomInts(r *rand.Rand) []int {
	arr := make([]int, r.IntN(100))
	for i := range arr {
		arr[i] = r.IntN(201) - 100
	}
	return arr
}

func randomFloats(r *rand.Rand) []float64 {
	arr := make([]float64, r.IntN(100))
	for i := range arr {
		arr[i] = r.Float64()*200 - 100
	}
	return arr
}

func randomStrings(r *rand.Rand) []string {
	arr := make([]string, r.IntN(20))
	for i := range arr {
		b := make([]byte, r.IntN(10))
		for j := range b {
			b[j] = byte('a' + r.IntN(26))
		}
		arr[i] = string(b)
	}
	return arr
}

func testSortRandom[T cmp.Ordered](t *testing.T, r *rand.Rand, gen func(*rand.Rand) []T) {
	for _, sc := range sortCases[T]() {
		t.Run(sc.name, func(t *testing.T) {
			for range 100 {
				arr := gen(r)
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
