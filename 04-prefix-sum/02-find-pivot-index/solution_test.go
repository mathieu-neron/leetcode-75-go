package solution

import (
	"testing"
)

func TestPivotIndex(t *testing.T) {
	t.Skip("TODO: implement pivotIndex")
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
	}
	for _, tc := range tests {
		got := pivotIndex(tc.nums)
		if got != tc.want {
			t.Errorf("pivotIndex(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}
