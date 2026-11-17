package sort

import (
	"cmp"
	"testing"
)

func insertionSortCases[T cmp.Ordered]() []sortImpl[T] {
	return []sortImpl[T]{
		{"classic", InsertionSortClassic[T]},
		{"shift", InsertionSortShift[T]},
		{"range", InsertionSortRange[T]},
		{"custom-binary", InsertionSortCustomBinary[T]},
		{"std-binary", InsertionSortStdBinary[T]},
	}
}

func TestInsertionCustomSort(t *testing.T) {
	testSortTableAll(t,
		insertionSortCases[int](),
		insertionSortCases[float64](),
		insertionSortCases[string](),
	)
}

func TestInsertionSortRandom(t *testing.T) {
	testSortRandomAll(t, newRand(),
		insertionSortCases[int](),
		insertionSortCases[float64](),
		insertionSortCases[string](),
	)
}

func FuzzInsertionSort(f *testing.F) {
	fuzzSort(f, insertionSortCases[int](), insertionSortCases[string]())
}

func BenchmarkInsertionSort(b *testing.B) {
	benchmarkSort(b, insertionSortCases[int]())
}
