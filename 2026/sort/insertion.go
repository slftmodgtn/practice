package sort

import (
	"cmp"
	"slices"
)

func InsertionSortClassic[T cmp.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func InsertionSortRange[T cmp.Ordered](arr []T) {
	for i, val := range arr {
		j := i - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = val
	}
}

// InsertionSortStdBinary сортирует срез с помощью бинарной сортировки вставками,
// используя стандартную функцию slices.BinarySearch для поиска позиции вставки.
// Это снижает количество сравнений элементов до O(N log N), хотя сдвиги остаются O(N^2).
func InsertionSortStdBinary[T cmp.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		// Сохраняем текущий элемент для вставки
		val := arr[i]

		// Ищем индекс для вставки в уже отсортированной части с помощью бинарного поиска
		// `slices.BinarySearch` возвращает индекс в диапазоне от 0 до len(arr) (включительно до len(arr):
		//
		//	Если число меньше всех элементов в отсортированной части, оно возвращает 0 (вставить в самое начало).
		//	Если число больше всех элементов, оно возвращает len(arr) (вставить в самый конец).
		//	Если число где-то посередине, возвращается точный индекс первого элемента, который больше или равен числу.
		j, _ := slices.BinarySearch(arr[:i], val)

		// Если элемент нужно вставить не на его текущую позицию
		if j < i {
			// Сдвигаем элементы вправо, освобождая место для вставки
			copy(arr[j+1:i+1], arr[j:i])
			// Вставляем сохраненный элемент на найденную позицию
			arr[j] = val
		}
	}
}

func InsertionSortCustomBinary[T cmp.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		low, high := 0, i
		for low < high {
			mid := low + (high-low)/2
			if arr[mid] > val {
				high = mid
			} else {
				low = mid + 1
			}
		}
		if low < i {
			copy(arr[low+1:i+1], arr[low:i])
			arr[low] = val
		}
	}
}