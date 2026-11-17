package sort

import (
	"cmp"
	"testing"
)

func selectionSortCases[T cmp.Ordered]() []sortImpl[T] {
	return []sortImpl[T]{
		{"selection", SelectionSort[T]},
	}
}

func TestSelectionCustomSort(t *testing.T) {
	testSortTableAll(t,
		selectionSortCases[int](),
		selectionSortCases[float64](),
		selectionSortCases[string](),
	)
}

func TestSelectionSortRandom(t *testing.T) {
	testSortRandomAll(t, newRand(),
		selectionSortCases[int](),
		selectionSortCases[float64](),
		selectionSortCases[string](),
	)
}

func FuzzSelectionSort(f *testing.F) {
	fuzzSort(f, selectionSortCases[int](), selectionSortCases[string]())
}

func BenchmarkSelectionSort(b *testing.B) {
	benchmarkSort(b, selectionSortCases[int]())
}
