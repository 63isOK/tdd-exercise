package peaks

import (
	"reflect"
	"testing"
)

func TestPrintPeaks(t *testing.T) {
	testCase := []struct {
		data        []int
		printString string
	}{
		{[]int{1, 2, 2, 2, 1}, "pos [1] peaks [2]"},
		{[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}, "pos [3 7] peaks [6 3]"},
		{[]int{1, 2, 2, 2, 3}, "pos [] peaks []"},
	}

	for _, x := range testCase {
		got := PrintPeaks(x.data)
		want := x.printString

		if got != want {
			t.Errorf("want '%s', got '%s', %v", want, got, x.data)
		}
	}
}

func TestFindPeaks(t *testing.T) {
	testCase := []struct {
		data []int
		pos  []int
	}{
		{[]int{1, 2, 2, 2, 1}, []int{1}},
		{[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}, []int{3, 7}},
		{[]int{1, 2, 2, 2, 3}, []int{}},
	}

	for _, x := range testCase {
		got := findPeaks(x.data)
		want := x.pos

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v, %v", want, got, x.data)
		}
	}
}
