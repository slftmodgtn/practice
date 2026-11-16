package array

// PushFront вставляет val в начало, вытесняя последний элемент.
// arr изменяется по месту через указатель.
func PushFront(arr [5]int, val int) (int, [5]int) {
	i := len(arr) - 1
	v := arr[i]
	for i > 0 {
		arr[i] = arr[i-1]
		i--
	}
	arr[0] = val

	return v, arr
}

// PushBack вставляет val в конец, вытесняя первый элемент.
// arr изменяется по месту через указатель.
func PushBack(arr [5]int, val int) (int, [5]int) {
	length, v := len(arr)-1, arr[0]
	for i := range length {
		arr[i] = arr[i+1]
	}
	arr[length] = val

	return v, arr
}
