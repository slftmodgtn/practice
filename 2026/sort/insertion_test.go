package main

import (
	"reflect"
	"testing"
)

type args struct {
	arr []int
}

type testDataStruct struct {
	name string
	args args
	want []int
}

func getTestData() []testDataStruct {
	return []testDataStruct{
		{"nil", args{[]int(nil)}, []int(nil)},
		{"empty", args{[]int{}}, []int{}},
		{"single", args{[]int{1}}, []int{1}},
		{"two", args{[]int{1, 2}}, []int{1, 2}},
		{"small", args{[]int{2, 1}}, []int{1, 2}},
		{"same", args{[]int{1, 1}}, []int{1, 1}},
		{"sorted", args{[]int{1, 2, 3, 4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"reversed", args{[]int{6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6}},
		{"doubled", args{[]int{3, 3, 2, 2, 1, 10, 0}}, []int{0, 1, 2, 2, 3, 3, 10}},
		{"random", args{[]int{10, 6, 2, 0, 34, 107, 36}}, []int{0, 2, 6, 10, 34, 36, 107}},
	}
}

func TestInsertionSortClassic(t *testing.T) {
	for _, tt := range getTestData() {
		t.Run(tt.name, func(t *testing.T) {
			if InsertionSortClassic(tt.args.arr); !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("InsertionSortClassic() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func BenchmarkInsertionSortClassic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSortClassic([]int{10, 6, 2, 0, 34, 107, 36})
	}
}

func TestInsertionSortRange(t *testing.T) {
	for _, tt := range getTestData() {
		t.Run(tt.name, func(t *testing.T) {
			if InsertionSortRange(tt.args.arr); !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("InsertionSortRange() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func BenchmarkInsertionSortRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSortRange([]int{10, 6, 2, 0, 34, 107, 36})
	}
}
