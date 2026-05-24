package solution

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	t.Skip("TODO: implement moveZeroes")
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}
	for _, tc := range tests {
		moveZeroes(tc.nums)
		if !reflect.DeepEqual(tc.nums, tc.want) {
			t.Errorf("after moveZeroes: got %v, want %v", tc.nums, tc.want)
		}
	}
}
