package move_elements

import (
	"testing"
)

const arrSize = 5

func TestArrShifting(t *testing.T) {
	testCases := []struct {
		name string
		fn   func([arrSize]int, int) (int, [arrSize]int)
	}{
		{"arrShiftCycleStraightforward", ArrShiftCycleStraightforward},
		{"arrShiftCyclingBackword", ArrShiftCyclingBackword},
	}

	type in struct {
		arr [arrSize]int
		val int
	}
	tests := []struct {
		name string
		in   in
	}{
		{"empty", in{[arrSize]int{}, 0}},
		{"single", in{[arrSize]int{1}, 6}},
		{"two", in{[arrSize]int{2, 1}, 6}},
		{"zeros before", in{[arrSize]int{0, 0, 0, 2, 1}, 6}},
	}
	for _, test := range tests {
		for _, testCase := range testCases {
			t.Run(testCase.name+"/"+test.name, func(t *testing.T) {
				wantArr := test.in.arr
				wantVal := test.in.val
				var got [arrSize]int
				for range arrSize + 1 {
					test.in.val, test.in.arr = testCase.fn(test.in.arr, test.in.val)
				}
				if test.in.arr != wantArr || test.in.val != wantVal {
					t.Errorf("%s(%v) = %v, want %v", test.name, test.in.arr, got, wantArr)
				}
			})
		}
	}
}
