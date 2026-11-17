package sort

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
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

var intCases = []testCase[int]{
	{"nil", []int(nil), []int(nil)},
	{"empty", []int{}, []int{}},
	{"single", []int{1}, []int{1}},
	{"small", []int{2, 1}, []int{1, 2}},
	{"same", []int{1, 1}, []int{1, 1}},
	{"sorted", []int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
	{"reversed", []int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	{"doubled", []int{3, 3, 2, 2, 1, 10, 0}, []int{0, 1, 2, 2, 3, 3, 10}},
	{"random", []int{10, 6, 2, 0, 34, 107, 36}, []int{0, 2, 6, 10, 34, 36, 107}},
}

var floatCases = []testCase[float64]{
	{"reversed", []float64{5.5, 4.4, 3.3, 2.2, 1.1}, []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
	{"same", []float64{1.5, 1.5, 1.5}, []float64{1.5, 1.5, 1.5}},
	{"random", []float64{10.5, -3.25, 0, 7.75, -100.5}, []float64{-100.5, -3.25, 0, 7.75, 10.5}},
}

var stringCases = []testCase[string]{
	{"reversed", []string{"delta", "charlie", "bravo", "alpha"}, []string{"alpha", "bravo", "charlie", "delta"}},
	{"same", []string{"x", "x"}, []string{"x", "x"}},
	{"case", []string{"banana", "Apple", "cherry"}, []string{"Apple", "banana", "cherry"}},
}

func testSortTable[T cmp.Ordered](t *testing.T, impls []sortImpl[T], cases []testCase[T]) {
	for _, sc := range impls {
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

func newRand() *rand.Rand {
	return rand.New(rand.NewPCG(1, 2))
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

func testSortRandom[T cmp.Ordered](t *testing.T, r *rand.Rand, impls []sortImpl[T], gen func(*rand.Rand) []T) {
	for _, sc := range impls {
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

func testSortTableAll(t *testing.T,
	intImpls []sortImpl[int],
	floatImpls []sortImpl[float64],
	stringImpls []sortImpl[string],
) {
	t.Run("int", func(t *testing.T) { testSortTable(t, intImpls, intCases) })
	t.Run("float64", func(t *testing.T) { testSortTable(t, floatImpls, floatCases) })
	t.Run("string", func(t *testing.T) { testSortTable(t, stringImpls, stringCases) })
}

func testSortRandomAll(t *testing.T, r *rand.Rand,
	intImpls []sortImpl[int],
	floatImpls []sortImpl[float64],
	stringImpls []sortImpl[string],
) {
	t.Run("int", func(t *testing.T) { testSortRandom(t, r, intImpls, randomInts) })
	t.Run("float64", func(t *testing.T) { testSortRandom(t, r, floatImpls, randomFloats) })
	t.Run("string", func(t *testing.T) { testSortRandom(t, r, stringImpls, randomStrings) })
}

func fuzzSort(f *testing.F, intImpls []sortImpl[int], stringImpls []sortImpl[string]) {
	f.Add([]byte{3, 1, 2}, "cba")
	f.Add([]byte{}, "")
	f.Fuzz(func(t *testing.T, data []byte, s string) {
		arr := make([]int, len(data))
		for i, b := range data {
			arr[i] = int(b) - 128
		}
		want := slices.Clone(arr)
		slices.Sort(want)
		for _, sc := range intImpls {
			got := slices.Clone(arr)
			sc.fn(got)
			if !slices.Equal(got, want) {
				t.Fatalf("%s(%v) = %v, want %v", sc.name, arr, got, want)
			}
		}

		strs := strings.Split(s, ",")
		wantStrs := slices.Clone(strs)
		slices.Sort(wantStrs)
		for _, sc := range stringImpls {
			got := slices.Clone(strs)
			sc.fn(got)
			if !slices.Equal(got, wantStrs) {
				t.Fatalf("%s(%q) = %q, want %q", sc.name, strs, got, wantStrs)
			}
		}
	})
}

const (
	benchKindRandom   = "random"
	benchKindSorted   = "sorted"
	benchKindReversed = "reversed"
)

var (
	benchSizes = []int{10, 100, 1000, 10000}
	benchKinds = []string{benchKindRandom, benchKindSorted, benchKindReversed}
)

func benchmarkData(size int, kind string) []int {
	r := rand.New(rand.NewPCG(uint64(size), 1))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = r.IntN(size * 10)
	}
	switch kind {
	case benchKindSorted:
		slices.Sort(arr)
	case benchKindReversed:
		slices.Sort(arr)
		slices.Reverse(arr)
	}
	return arr
}

func benchmarkSort(b *testing.B, impls []sortImpl[int]) {
	// стд алгоритм сортировки для сравнения скорости выполнения c реализованными
	impls = append(impls, sortImpl[int]{"stdlib", func(arr []int) { slices.Sort(arr) }})

	for _, size := range benchSizes {
		for _, kind := range benchKinds {
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
