package move_elements

// ArrShiftCycleStraightforward добавление val в начало массива со смещением. in-place
// возвращает обновленный массив и смещенное значение
func ArrShiftCycleStraightforward(arr [5]int, val int) (int, [5]int) {
	v := arr[len(arr) - 1]
	// смещение элементов
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = val
	
	return v, arr
}

// ArrShiftCyclingBackword добавление val в конец массива со смещением. in-place
// возвращает обновленный массив и смещенное значение
func ArrShiftCyclingBackword(arr [5]int, val int) (int, [5]int) {
	v := arr[0]
	// смещение элементов
	for i := range len(arr) - 1 {
		arr[i] = arr[i+1]
	}
	arr[len(arr) - 1] = val

	return v, arr
}
